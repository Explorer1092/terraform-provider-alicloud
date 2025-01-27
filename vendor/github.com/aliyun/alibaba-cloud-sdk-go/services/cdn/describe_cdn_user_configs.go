package cdn

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

// DescribeCdnUserConfigs invokes the cdn.DescribeCdnUserConfigs API synchronously
func (client *Client) DescribeCdnUserConfigs(request *DescribeCdnUserConfigsRequest) (response *DescribeCdnUserConfigsResponse, err error) {
	response = CreateDescribeCdnUserConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnUserConfigsWithChan invokes the cdn.DescribeCdnUserConfigs API asynchronously
func (client *Client) DescribeCdnUserConfigsWithChan(request *DescribeCdnUserConfigsRequest) (<-chan *DescribeCdnUserConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnUserConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnUserConfigs(request)
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

// DescribeCdnUserConfigsWithCallback invokes the cdn.DescribeCdnUserConfigs API asynchronously
func (client *Client) DescribeCdnUserConfigsWithCallback(request *DescribeCdnUserConfigsRequest, callback func(response *DescribeCdnUserConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnUserConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnUserConfigs(request)
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

// DescribeCdnUserConfigsRequest is the request struct for api DescribeCdnUserConfigs
type DescribeCdnUserConfigsRequest struct {
	*requests.RpcRequest
	FunctionName string `position:"Query" name:"FunctionName"`
}

// DescribeCdnUserConfigsResponse is the response struct for api DescribeCdnUserConfigs
type DescribeCdnUserConfigsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Configs   []Config `json:"Configs" xml:"Configs"`
}

// CreateDescribeCdnUserConfigsRequest creates a request to invoke DescribeCdnUserConfigs API
func CreateDescribeCdnUserConfigsRequest() (request *DescribeCdnUserConfigsRequest) {
	request = &DescribeCdnUserConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnUserConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnUserConfigsResponse creates a response to parse from DescribeCdnUserConfigs response
func CreateDescribeCdnUserConfigsResponse() (response *DescribeCdnUserConfigsResponse) {
	response = &DescribeCdnUserConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
