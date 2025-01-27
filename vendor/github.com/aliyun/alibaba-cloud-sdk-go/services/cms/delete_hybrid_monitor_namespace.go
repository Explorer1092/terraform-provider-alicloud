package cms

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

// DeleteHybridMonitorNamespace invokes the cms.DeleteHybridMonitorNamespace API synchronously
func (client *Client) DeleteHybridMonitorNamespace(request *DeleteHybridMonitorNamespaceRequest) (response *DeleteHybridMonitorNamespaceResponse, err error) {
	response = CreateDeleteHybridMonitorNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteHybridMonitorNamespaceWithChan invokes the cms.DeleteHybridMonitorNamespace API asynchronously
func (client *Client) DeleteHybridMonitorNamespaceWithChan(request *DeleteHybridMonitorNamespaceRequest) (<-chan *DeleteHybridMonitorNamespaceResponse, <-chan error) {
	responseChan := make(chan *DeleteHybridMonitorNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteHybridMonitorNamespace(request)
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

// DeleteHybridMonitorNamespaceWithCallback invokes the cms.DeleteHybridMonitorNamespace API asynchronously
func (client *Client) DeleteHybridMonitorNamespaceWithCallback(request *DeleteHybridMonitorNamespaceRequest, callback func(response *DeleteHybridMonitorNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteHybridMonitorNamespaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteHybridMonitorNamespace(request)
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

// DeleteHybridMonitorNamespaceRequest is the request struct for api DeleteHybridMonitorNamespace
type DeleteHybridMonitorNamespaceRequest struct {
	*requests.RpcRequest
	Namespace string           `position:"Query" name:"Namespace"`
	ClearFlag requests.Boolean `position:"Query" name:"ClearFlag"`
}

// DeleteHybridMonitorNamespaceResponse is the response struct for api DeleteHybridMonitorNamespace
type DeleteHybridMonitorNamespaceResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateDeleteHybridMonitorNamespaceRequest creates a request to invoke DeleteHybridMonitorNamespace API
func CreateDeleteHybridMonitorNamespaceRequest() (request *DeleteHybridMonitorNamespaceRequest) {
	request = &DeleteHybridMonitorNamespaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteHybridMonitorNamespace", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteHybridMonitorNamespaceResponse creates a response to parse from DeleteHybridMonitorNamespace response
func CreateDeleteHybridMonitorNamespaceResponse() (response *DeleteHybridMonitorNamespaceResponse) {
	response = &DeleteHybridMonitorNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
