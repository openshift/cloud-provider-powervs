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
	"log"
	"testing"

	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/stretchr/testify/assert"
)

func TestPopulatePowerVSNodeMetadata(t *testing.T) {
	powerVSClient := ibmPowerVSClient{
		provider:        Provider{PowerVSRegion: "lon", PowerVSZone: "lon4"},
		sdk:             &powerVSTestClient{},
		cloudInstanceID: "testCloudInstanceId",
	}
	newNode := NodeMetadata{}
	err := powerVSClient.populateNodeMetadata("testNode", &newNode)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "192.168.1.1", newNode.InternalIP, "Unexpected InternalIP")
	assert.Equal(t, "740aabcb-5096-4024-4324-12327e8d0def", newNode.WorkerID, "Unexpected WorkerID")
	assert.Equal(t, "s922", newNode.InstanceType, "Unexpected InstanceType")
	assert.Equal(t, "lon4", newNode.FailureDomain, "Unexpected FailureDomain")
	assert.Equal(t, "lon", newNode.Region, "Unexpected Region")
}

type powerVSTestClient struct {
}

func (p *powerVSTestClient) GetInstanceByName(name string) (*models.PVMInstance, error) {
	pvmInstanceID := "740aabcb-5096-4024-4324-12327e8d0def"
	return &models.PVMInstance{
		PvmInstanceID: &pvmInstanceID,
		SysType:       "s922",
		Networks: []*models.PVMInstanceNetwork{
			{IPAddress: "192.168.1.1"},
		},
	}, nil

}

func (p *powerVSTestClient) GetInstances() (*models.PVMInstances, error) {
	pvmInstanceID := "740aabcb-5096-4024-4324-12327e8d0def"
	serverName := "TestMachine"
	return &models.PVMInstances{PvmInstances: []*models.PVMInstanceReference{
		{
			Networks: []*models.PVMInstanceNetwork{
				{IPAddress: "192.168.1.1"},
			},
			PvmInstanceID: &pvmInstanceID,
			ServerName:    &serverName,
			SysType:       "s922",
		},
	}}, nil
}
