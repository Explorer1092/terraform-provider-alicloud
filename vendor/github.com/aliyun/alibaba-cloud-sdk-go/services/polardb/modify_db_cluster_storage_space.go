package polardb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDBClusterStorageSpace invokes the polardb.ModifyDBClusterStorageSpace API synchronously
func (client *Client) ModifyDBClusterStorageSpace(request *ModifyDBClusterStorageSpaceRequest) (response *ModifyDBClusterStorageSpaceResponse, err error) {
	response = CreateModifyDBClusterStorageSpaceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterStorageSpaceWithChan invokes the polardb.ModifyDBClusterStorageSpace API asynchronously
func (client *Client) ModifyDBClusterStorageSpaceWithChan(request *ModifyDBClusterStorageSpaceRequest) (<-chan *ModifyDBClusterStorageSpaceResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterStorageSpaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterStorageSpace(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDBClusterStorageSpaceWithCallback invokes the polardb.ModifyDBClusterStorageSpace API asynchronously
func (client *Client) ModifyDBClusterStorageSpaceWithCallback(request *ModifyDBClusterStorageSpaceRequest, callback func(response *ModifyDBClusterStorageSpaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterStorageSpaceResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterStorageSpace(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDBClusterStorageSpaceRequest is the request struct for api ModifyDBClusterStorageSpace
type ModifyDBClusterStorageSpaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PlannedEndTime       string           `position:"Query" name:"PlannedEndTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PlannedStartTime     string           `position:"Query" name:"PlannedStartTime"`
	SubCategory          string           `position:"Query" name:"SubCategory"`
	StorageSpace         requests.Integer `position:"Query" name:"StorageSpace"`
}

// ModifyDBClusterStorageSpaceResponse is the response struct for api ModifyDBClusterStorageSpace
type ModifyDBClusterStorageSpaceResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
	OrderId     string `json:"OrderId" xml:"OrderId"`
}

// CreateModifyDBClusterStorageSpaceRequest creates a request to invoke ModifyDBClusterStorageSpace API
func CreateModifyDBClusterStorageSpaceRequest() (request *ModifyDBClusterStorageSpaceRequest) {
	request = &ModifyDBClusterStorageSpaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterStorageSpace", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterStorageSpaceResponse creates a response to parse from ModifyDBClusterStorageSpace response
func CreateModifyDBClusterStorageSpaceResponse() (response *ModifyDBClusterStorageSpaceResponse) {
	response = &ModifyDBClusterStorageSpaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}