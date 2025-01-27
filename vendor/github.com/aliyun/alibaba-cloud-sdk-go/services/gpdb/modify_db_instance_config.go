package gpdb

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

// ModifyDBInstanceConfig invokes the gpdb.ModifyDBInstanceConfig API synchronously
func (client *Client) ModifyDBInstanceConfig(request *ModifyDBInstanceConfigRequest) (response *ModifyDBInstanceConfigResponse, err error) {
	response = CreateModifyDBInstanceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceConfigWithChan invokes the gpdb.ModifyDBInstanceConfig API asynchronously
func (client *Client) ModifyDBInstanceConfigWithChan(request *ModifyDBInstanceConfigRequest) (<-chan *ModifyDBInstanceConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceConfig(request)
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

// ModifyDBInstanceConfigWithCallback invokes the gpdb.ModifyDBInstanceConfig API asynchronously
func (client *Client) ModifyDBInstanceConfigWithCallback(request *ModifyDBInstanceConfigRequest, callback func(response *ModifyDBInstanceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceConfig(request)
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

// ModifyDBInstanceConfigRequest is the request struct for api ModifyDBInstanceConfig
type ModifyDBInstanceConfigRequest struct {
	*requests.RpcRequest
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	ServerlessResource    requests.Integer `position:"Query" name:"ServerlessResource"`
	IdleTime              requests.Integer `position:"Query" name:"IdleTime"`
}

// ModifyDBInstanceConfigResponse is the response struct for api ModifyDBInstanceConfig
type ModifyDBInstanceConfigResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	DbInstanceId string `json:"DbInstanceId" xml:"DbInstanceId"`
	Status       bool   `json:"Status" xml:"Status"`
}

// CreateModifyDBInstanceConfigRequest creates a request to invoke ModifyDBInstanceConfig API
func CreateModifyDBInstanceConfigRequest() (request *ModifyDBInstanceConfigRequest) {
	request = &ModifyDBInstanceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ModifyDBInstanceConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceConfigResponse creates a response to parse from ModifyDBInstanceConfig response
func CreateModifyDBInstanceConfigResponse() (response *ModifyDBInstanceConfigResponse) {
	response = &ModifyDBInstanceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
