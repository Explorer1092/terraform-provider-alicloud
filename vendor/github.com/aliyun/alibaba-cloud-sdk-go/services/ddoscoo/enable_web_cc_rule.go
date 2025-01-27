package ddoscoo

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

// EnableWebCCRule invokes the ddoscoo.EnableWebCCRule API synchronously
func (client *Client) EnableWebCCRule(request *EnableWebCCRuleRequest) (response *EnableWebCCRuleResponse, err error) {
	response = CreateEnableWebCCRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableWebCCRuleWithChan invokes the ddoscoo.EnableWebCCRule API asynchronously
func (client *Client) EnableWebCCRuleWithChan(request *EnableWebCCRuleRequest) (<-chan *EnableWebCCRuleResponse, <-chan error) {
	responseChan := make(chan *EnableWebCCRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableWebCCRule(request)
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

// EnableWebCCRuleWithCallback invokes the ddoscoo.EnableWebCCRule API asynchronously
func (client *Client) EnableWebCCRuleWithCallback(request *EnableWebCCRuleRequest, callback func(response *EnableWebCCRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableWebCCRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableWebCCRule(request)
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

// EnableWebCCRuleRequest is the request struct for api EnableWebCCRule
type EnableWebCCRuleRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
}

// EnableWebCCRuleResponse is the response struct for api EnableWebCCRule
type EnableWebCCRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableWebCCRuleRequest creates a request to invoke EnableWebCCRule API
func CreateEnableWebCCRuleRequest() (request *EnableWebCCRuleRequest) {
	request = &EnableWebCCRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "EnableWebCCRule", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableWebCCRuleResponse creates a response to parse from EnableWebCCRule response
func CreateEnableWebCCRuleResponse() (response *EnableWebCCRuleResponse) {
	response = &EnableWebCCRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
