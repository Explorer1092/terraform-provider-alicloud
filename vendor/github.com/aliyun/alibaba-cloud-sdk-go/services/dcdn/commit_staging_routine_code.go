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

// CommitStagingRoutineCode invokes the dcdn.CommitStagingRoutineCode API synchronously
func (client *Client) CommitStagingRoutineCode(request *CommitStagingRoutineCodeRequest) (response *CommitStagingRoutineCodeResponse, err error) {
	response = CreateCommitStagingRoutineCodeResponse()
	err = client.DoAction(request, response)
	return
}

// CommitStagingRoutineCodeWithChan invokes the dcdn.CommitStagingRoutineCode API asynchronously
func (client *Client) CommitStagingRoutineCodeWithChan(request *CommitStagingRoutineCodeRequest) (<-chan *CommitStagingRoutineCodeResponse, <-chan error) {
	responseChan := make(chan *CommitStagingRoutineCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CommitStagingRoutineCode(request)
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

// CommitStagingRoutineCodeWithCallback invokes the dcdn.CommitStagingRoutineCode API asynchronously
func (client *Client) CommitStagingRoutineCodeWithCallback(request *CommitStagingRoutineCodeRequest, callback func(response *CommitStagingRoutineCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CommitStagingRoutineCodeResponse
		var err error
		defer close(result)
		response, err = client.CommitStagingRoutineCode(request)
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

// CommitStagingRoutineCodeRequest is the request struct for api CommitStagingRoutineCode
type CommitStagingRoutineCodeRequest struct {
	*requests.RpcRequest
	CodeDescription string `position:"Body" name:"CodeDescription"`
	Name            string `position:"Body" name:"Name"`
}

// CommitStagingRoutineCodeResponse is the response struct for api CommitStagingRoutineCode
type CommitStagingRoutineCodeResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateCommitStagingRoutineCodeRequest creates a request to invoke CommitStagingRoutineCode API
func CreateCommitStagingRoutineCodeRequest() (request *CommitStagingRoutineCodeRequest) {
	request = &CommitStagingRoutineCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "CommitStagingRoutineCode", "", "")
	request.Method = requests.POST
	return
}

// CreateCommitStagingRoutineCodeResponse creates a response to parse from CommitStagingRoutineCode response
func CreateCommitStagingRoutineCodeResponse() (response *CommitStagingRoutineCodeResponse) {
	response = &CommitStagingRoutineCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
