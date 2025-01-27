package adb

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

// UnbindDBResourceGroupWithUser invokes the adb.UnbindDBResourceGroupWithUser API synchronously
func (client *Client) UnbindDBResourceGroupWithUser(request *UnbindDBResourceGroupWithUserRequest) (response *UnbindDBResourceGroupWithUserResponse, err error) {
	response = CreateUnbindDBResourceGroupWithUserResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindDBResourceGroupWithUserWithChan invokes the adb.UnbindDBResourceGroupWithUser API asynchronously
func (client *Client) UnbindDBResourceGroupWithUserWithChan(request *UnbindDBResourceGroupWithUserRequest) (<-chan *UnbindDBResourceGroupWithUserResponse, <-chan error) {
	responseChan := make(chan *UnbindDBResourceGroupWithUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindDBResourceGroupWithUser(request)
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

// UnbindDBResourceGroupWithUserWithCallback invokes the adb.UnbindDBResourceGroupWithUser API asynchronously
func (client *Client) UnbindDBResourceGroupWithUserWithCallback(request *UnbindDBResourceGroupWithUserRequest, callback func(response *UnbindDBResourceGroupWithUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindDBResourceGroupWithUserResponse
		var err error
		defer close(result)
		response, err = client.UnbindDBResourceGroupWithUser(request)
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

// UnbindDBResourceGroupWithUserRequest is the request struct for api UnbindDBResourceGroupWithUser
type UnbindDBResourceGroupWithUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	GroupUser            string           `position:"Query" name:"GroupUser"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	GroupName            string           `position:"Query" name:"GroupName"`
}

// UnbindDBResourceGroupWithUserResponse is the response struct for api UnbindDBResourceGroupWithUser
type UnbindDBResourceGroupWithUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindDBResourceGroupWithUserRequest creates a request to invoke UnbindDBResourceGroupWithUser API
func CreateUnbindDBResourceGroupWithUserRequest() (request *UnbindDBResourceGroupWithUserRequest) {
	request = &UnbindDBResourceGroupWithUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "UnbindDBResourceGroupWithUser", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindDBResourceGroupWithUserResponse creates a response to parse from UnbindDBResourceGroupWithUser response
func CreateUnbindDBResourceGroupWithUserResponse() (response *UnbindDBResourceGroupWithUserResponse) {
	response = &UnbindDBResourceGroupWithUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
