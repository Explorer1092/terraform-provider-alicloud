package dcdn

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

// CreateDcdnSubTask invokes the dcdn.CreateDcdnSubTask API synchronously
func (client *Client) CreateDcdnSubTask(request *CreateDcdnSubTaskRequest) (response *CreateDcdnSubTaskResponse, err error) {
	response = CreateCreateDcdnSubTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDcdnSubTaskWithChan invokes the dcdn.CreateDcdnSubTask API asynchronously
func (client *Client) CreateDcdnSubTaskWithChan(request *CreateDcdnSubTaskRequest) (<-chan *CreateDcdnSubTaskResponse, <-chan error) {
	responseChan := make(chan *CreateDcdnSubTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDcdnSubTask(request)
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

// CreateDcdnSubTaskWithCallback invokes the dcdn.CreateDcdnSubTask API asynchronously
func (client *Client) CreateDcdnSubTaskWithCallback(request *CreateDcdnSubTaskRequest, callback func(response *CreateDcdnSubTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDcdnSubTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateDcdnSubTask(request)
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

// CreateDcdnSubTaskRequest is the request struct for api CreateDcdnSubTask
type CreateDcdnSubTaskRequest struct {
	*requests.RpcRequest
	ReportIds  string `position:"Body" name:"ReportIds"`
	DomainName string `position:"Body" name:"DomainName"`
}

// CreateDcdnSubTaskResponse is the response struct for api CreateDcdnSubTask
type CreateDcdnSubTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDcdnSubTaskRequest creates a request to invoke CreateDcdnSubTask API
func CreateCreateDcdnSubTaskRequest() (request *CreateDcdnSubTaskRequest) {
	request = &CreateDcdnSubTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "CreateDcdnSubTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDcdnSubTaskResponse creates a response to parse from CreateDcdnSubTask response
func CreateCreateDcdnSubTaskResponse() (response *CreateDcdnSubTaskResponse) {
	response = &CreateDcdnSubTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
