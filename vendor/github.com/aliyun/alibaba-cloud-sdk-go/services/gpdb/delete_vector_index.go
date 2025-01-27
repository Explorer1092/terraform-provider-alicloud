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

// DeleteVectorIndex invokes the gpdb.DeleteVectorIndex API synchronously
func (client *Client) DeleteVectorIndex(request *DeleteVectorIndexRequest) (response *DeleteVectorIndexResponse, err error) {
	response = CreateDeleteVectorIndexResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVectorIndexWithChan invokes the gpdb.DeleteVectorIndex API asynchronously
func (client *Client) DeleteVectorIndexWithChan(request *DeleteVectorIndexRequest) (<-chan *DeleteVectorIndexResponse, <-chan error) {
	responseChan := make(chan *DeleteVectorIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVectorIndex(request)
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

// DeleteVectorIndexWithCallback invokes the gpdb.DeleteVectorIndex API asynchronously
func (client *Client) DeleteVectorIndexWithCallback(request *DeleteVectorIndexRequest, callback func(response *DeleteVectorIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVectorIndexResponse
		var err error
		defer close(result)
		response, err = client.DeleteVectorIndex(request)
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

// DeleteVectorIndexRequest is the request struct for api DeleteVectorIndex
type DeleteVectorIndexRequest struct {
	*requests.RpcRequest
	ManagerAccount         string           `position:"Query" name:"ManagerAccount"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ManagerAccountPassword string           `position:"Query" name:"ManagerAccountPassword"`
	Collection             string           `position:"Query" name:"Collection"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	Namespace              string           `position:"Query" name:"Namespace"`
}

// DeleteVectorIndexResponse is the response struct for api DeleteVectorIndex
type DeleteVectorIndexResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateDeleteVectorIndexRequest creates a request to invoke DeleteVectorIndex API
func CreateDeleteVectorIndexRequest() (request *DeleteVectorIndexRequest) {
	request = &DeleteVectorIndexRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DeleteVectorIndex", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteVectorIndexResponse creates a response to parse from DeleteVectorIndex response
func CreateDeleteVectorIndexResponse() (response *DeleteVectorIndexResponse) {
	response = &DeleteVectorIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
