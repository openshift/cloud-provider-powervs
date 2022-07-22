/*******************************************************************************
* IBM Cloud Kubernetes Service, 5737-D43
* (C) Copyright IBM Corp. 2022 All Rights Reserved.
*
* SPDX-License-Identifier: Apache2.0
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*******************************************************************************/

package ibm

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v5/core"

	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

// dhcpCacheStore is a cache store to hold the Power VS VM DHCP IP.
var dhcpCacheStore cache.Store

func init() {
	dhcpCacheStore = initialiseDHCPCacheStore()
}

// Client is a wrapper object for actual PowerVS SDK clients to allow for easier testing.
type Client interface {
	GetInstances() (*models.PVMInstances, error)
	GetInstanceByName(name string) (*models.PVMInstance, error)
	GetDHCPServers() (models.DHCPServers, error)
	GetDHCPServerByID(string) (*models.DHCPServerDetail, error)
}

// ibmPowerVSClient makes call to IBM Cloud Power VS APIs
type ibmPowerVSClient struct {
	provider        Provider
	sdk             Client
	cloudInstanceID string
}

// powerVSClient helps in initializing ibmPowerVSClient Client and implements Client interface
type powerVSClient struct {
	cloudInstanceID string
	instanceClient  *instance.IBMPIInstanceClient
	dhcpClient      *instance.IBMPIDhcpClient
}

// newPowerVSSdkClient initializes a new sdk client and can be overridden by testing
var newPowerVSSdkClient = func(provider Provider) (Client, error) {
	credential, err := readCredential(provider)
	if err != nil {
		klog.Errorf("Failed to read the credentials, Error: %v", err)
		return nil, err
	}

	// Create the authenticator
	authenticator := &core.IamAuthenticator{
		ApiKey: credential,
	}

	// Create the session options struct
	options := &ibmpisession.IBMPIOptions{
		Authenticator: authenticator,
		UserAccount:   provider.AccountID,
		Region:        provider.PowerVSRegion,
		Zone:          provider.PowerVSZone,
	}

	// Construct the session service instance
	session, err := ibmpisession.NewIBMPISession(options)
	if err != nil {
		klog.Errorf("Failed to create new IBMPISession, Error: %v", err)
		return nil, err
	}

	client := &powerVSClient{
		cloudInstanceID: provider.PowerVSCloudInstanceID,
	}
	ctx := context.Background()
	client.instanceClient = instance.NewIBMPIInstanceClient(ctx, session, client.cloudInstanceID)
	client.dhcpClient = instance.NewIBMPIDhcpClient(ctx, session, client.cloudInstanceID)
	return client, nil
}

// newPowerVSClient initializes a new validated powerVSClient
func newPowerVSClient(provider Provider) (*ibmPowerVSClient, error) {
	// create Power VS sdk client
	sdk, err := newPowerVSSdkClient(provider)
	if err != nil {
		klog.Errorf("Failed to create newPowerVSSdkClient, Error: %v", err)
		return nil, err
	}

	return &ibmPowerVSClient{
		provider: provider,
		sdk:      sdk,
	}, nil
}

// isProviderPowerVS returns true when PowerVS specific parameters are set in Provider
func isProviderPowerVS(provider Provider) bool {
	if provider.PowerVSCloudInstanceID != "" && provider.PowerVSRegion != "" && provider.PowerVSZone != "" {
		return true
	}
	return false
}

// GetInstanceByName return the Power VS instance corresponding to the provided name
func (p *powerVSClient) GetInstanceByName(name string) (*models.PVMInstance, error) {
	instances, err := p.GetInstances()
	if err != nil {
		klog.Errorf("Failed to get instance list, Error: %v", err)
		return nil, fmt.Errorf("failed to get the instance list, Error: %v", err)
	}
	for _, i := range instances.PvmInstances {
		if *i.ServerName == name {
			klog.Infof("instance name: %s id %s", name, *i.PvmInstanceID)
			return p.instanceClient.Get(*i.PvmInstanceID)
		}
	}
	return nil, fmt.Errorf("Instance %s not found in %s cloud instance ", name, p.cloudInstanceID)
}

// GetInstances returns all the Power VS instances
func (p *powerVSClient) GetInstances() (*models.PVMInstances, error) {
	return p.instanceClient.GetAll()
}

func (p *powerVSClient) GetDHCPServers() (models.DHCPServers, error) {
	return p.dhcpClient.GetAll()
}

func (p *powerVSClient) GetDHCPServerByID(id string) (*models.DHCPServerDetail, error) {
	return p.dhcpClient.Get(id)
}

// populateNodeMetadata forms the node metadata from instance details
func (p *ibmPowerVSClient) populateNodeMetadata(nodeName string, node *NodeMetadata) error {

	// Try to fetch the nodeMetadata from cache
	obj, exists, err := dhcpCacheStore.GetByKey(nodeName)
	if err != nil {
		klog.Errorf("Node %s failed to fetch the node metadata from cache, error: %v", nodeName, err)
	}
	if exists {
		klog.Infof("Node %s found metadata %+v from DHCP cache", nodeName, obj.(nodeMetadataCache).Metadata)
		node = obj.(nodeMetadataCache).Metadata
		return nil
	}

	// Get Power VS Instance.
	pvsInstance, err := p.sdk.GetInstanceByName(nodeName)
	if err != nil {
		return err
	}
	// Check if instance is not nil.
	if pvsInstance == nil {
		return errors.New("Could not retrieve a Power instance: name=" + nodeName)
	}
	node.WorkerID = *pvsInstance.PvmInstanceID
	klog.Infof("Node %s worker id is %s", nodeName, node.WorkerID)

	node.InstanceType = pvsInstance.SysType
	klog.Infof("Node %s instance type is %s", nodeName, node.InstanceType)

	node.Region = p.provider.PowerVSRegion
	klog.Infof("Node %s region is %s", nodeName, node.Region)

	node.FailureDomain = p.provider.PowerVSZone
	klog.Infof("Node %s failureDomain is %s", nodeName, node.FailureDomain)

	for _, network := range pvsInstance.Networks {
		if strings.TrimSpace(network.ExternalIP) != "" {
			node.ExternalIP = strings.TrimSpace(network.ExternalIP)
		}
		if strings.TrimSpace(network.IPAddress) != "" {
			node.InternalIP = strings.TrimSpace(network.IPAddress)
		}
	}

	if node.ExternalIP == "" && node.InternalIP == "" {
		// If node ExternalIP and InternalIP is empty, try to fetch the IP from dhcp server.
		klog.Infof("Node %s fetching IP from DHCP server", nodeName)
		// Fetch the Network attached to instance.
		network, err := getPowerVSNetwork(pvsInstance)
		if err != nil {
			klog.Errorf("failed to fetch Power VS Network name error: %v", err)
			return err
		}
		// for DHCP network type will be "dynamic" for other networks type will be "fixed"
		if network.Type != "dynamic" {
			errStr := fmt.Errorf("Node %s attached with network %s of type %s expecting Network Type to be dynamic to fetch IP from DHCP server ", nodeName, network.NetworkName, network.Type)
			klog.Error(errStr.Error())
			return errStr
		}
		// Fetch the DHCP server ID.
		dhcpServerID, err := p.getDHCPServerID(network.NetworkName)
		if err != nil {
			klog.Errorf("failed to fetch dhcp server id error: %v", err)
			return err
		}
		dhcpServerDetails, err := p.sdk.GetDHCPServerByID(dhcpServerID)
		if err != nil {
			klog.Errorf("failed to fetch dhcp server details with id: %s error: %v", dhcpServerID, err)
			return err
		}
		for _, lease := range dhcpServerDetails.Leases {
			if network.MacAddress == *lease.InstanceMacAddress {
				node.InternalIP = *lease.InstanceIP
				klog.Infof("Node %s found internal ip %s from dhcp lease", nodeName, node.InternalIP)
			}
		}
		if len(node.InternalIP) == 0 {
			errStr := fmt.Errorf("not able to find internal ip for pvm instance: %s with attached network name: %s", nodeName, network.NetworkName)
			klog.Error(errStr)
			return errStr
		}
	}

	// Update the cache with the node metadata
	err = dhcpCacheStore.Add(nodeMetadataCache{
		Name:     nodeName,
		Metadata: node,
	})
	if err != nil {
		klog.Errorf("Nod %s failed to add node metadata to cache, Error %v", err)
	}
	klog.Infof("Node %s internal IP is %s", nodeName, node.InternalIP)
	klog.Infof("Node %s external IP is %s", nodeName, node.ExternalIP)
	return nil
}

// getDHCPServerID fetches and returns the DHCP server ID.
func (p *ibmPowerVSClient) getDHCPServerID(networkName string) (string, error) {
	// If the DHCP server ID is not provided try to fetch it.
	// Get all the DHCP servers.
	dhcpServer, err := p.sdk.GetDHCPServers()
	if err != nil {
		klog.Errorf("failed to get DHCP server error: %v", err)
		return "", err
	}
	// Get the DHCP server ID.
	for _, server := range dhcpServer {
		if *server.Network.Name == networkName {
			klog.Infof("Found DHCP server with ID %s for network %s", *server.ID, networkName)
			return *server.ID, nil
		}
	}
	return "", fmt.Errorf("not able to get DHCP server ID for network %s", networkName)
}

// getPowerVSNetwork fetches and returns the Power VS Network
func getPowerVSNetwork(instance *models.PVMInstance) (*models.PVMInstanceNetwork, error) {
	// currently its assumed that there will be only 1 network attached to instance
	if len(instance.Networks) != 1 {
		errStr := fmt.Errorf("expecting only one network to be attached to vm %s but got %d", *instance.ServerName, len(instance.Networks))
		klog.Error(errStr.Error())
		return nil, errStr
	}
	return instance.Networks[0], nil
}
