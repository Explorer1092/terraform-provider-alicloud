package cloudapi

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

// AttachPlugin invokes the cloudapi.AttachPlugin API synchronously
func (client *Client) AttachPlugin(request *AttachPluginRequest) (response *AttachPluginResponse, err error) {
	response = CreateAttachPluginResponse()
	err = client.DoAction(request, response)
	return
}

// AttachPluginWithChan invokes the cloudapi.AttachPlugin API asynchronously
func (client *Client) AttachPluginWithChan(request *AttachPluginRequest) (<-chan *AttachPluginResponse, <-chan error) {
	responseChan := make(chan *AttachPluginResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachPlugin(request)
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

// AttachPluginWithCallback invokes the cloudapi.AttachPlugin API asynchronously
func (client *Client) AttachPluginWithCallback(request *AttachPluginRequest, callback func(response *AttachPluginResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachPluginResponse
		var err error
		defer close(result)
		response, err = client.AttachPlugin(request)
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

// AttachPluginRequest is the request struct for api AttachPlugin
type AttachPluginRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	PluginId      string `position:"Query" name:"PluginId"`
	GroupId       string `position:"Query" name:"GroupId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiId         string `position:"Query" name:"ApiId"`
	ApiIds        string `position:"Query" name:"ApiIds"`
}

// AttachPluginResponse is the response struct for api AttachPlugin
type AttachPluginResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachPluginRequest creates a request to invoke AttachPlugin API
func CreateAttachPluginRequest() (request *AttachPluginRequest) {
	request = &AttachPluginRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "AttachPlugin", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachPluginResponse creates a response to parse from AttachPlugin response
func CreateAttachPluginResponse() (response *AttachPluginResponse) {
	response = &AttachPluginResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
