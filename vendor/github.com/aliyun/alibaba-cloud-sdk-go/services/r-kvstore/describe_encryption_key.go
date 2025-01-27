package r_kvstore

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

// DescribeEncryptionKey invokes the r_kvstore.DescribeEncryptionKey API synchronously
func (client *Client) DescribeEncryptionKey(request *DescribeEncryptionKeyRequest) (response *DescribeEncryptionKeyResponse, err error) {
	response = CreateDescribeEncryptionKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEncryptionKeyWithChan invokes the r_kvstore.DescribeEncryptionKey API asynchronously
func (client *Client) DescribeEncryptionKeyWithChan(request *DescribeEncryptionKeyRequest) (<-chan *DescribeEncryptionKeyResponse, <-chan error) {
	responseChan := make(chan *DescribeEncryptionKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEncryptionKey(request)
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

// DescribeEncryptionKeyWithCallback invokes the r_kvstore.DescribeEncryptionKey API asynchronously
func (client *Client) DescribeEncryptionKeyWithCallback(request *DescribeEncryptionKeyRequest, callback func(response *DescribeEncryptionKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEncryptionKeyResponse
		var err error
		defer close(result)
		response, err = client.DescribeEncryptionKey(request)
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

// DescribeEncryptionKeyRequest is the request struct for api DescribeEncryptionKey
type DescribeEncryptionKeyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EncryptionKey        string           `position:"Query" name:"EncryptionKey"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeEncryptionKeyResponse is the response struct for api DescribeEncryptionKey
type DescribeEncryptionKeyResponse struct {
	*responses.BaseResponse
	DeleteDate          string `json:"DeleteDate" xml:"DeleteDate"`
	RequestId           string `json:"RequestId" xml:"RequestId"`
	Description         string `json:"Description" xml:"Description"`
	Origin              string `json:"Origin" xml:"Origin"`
	MaterialExpireTime  string `json:"MaterialExpireTime" xml:"MaterialExpireTime"`
	EncryptionKeyStatus string `json:"EncryptionKeyStatus" xml:"EncryptionKeyStatus"`
	KeyUsage            string `json:"KeyUsage" xml:"KeyUsage"`
	EncryptionKey       string `json:"EncryptionKey" xml:"EncryptionKey"`
	Creator             string `json:"Creator" xml:"Creator"`
	EncryptionName      string `json:"EncryptionName" xml:"EncryptionName"`
	RoleArn             string `json:"RoleArn" xml:"RoleArn"`
}

// CreateDescribeEncryptionKeyRequest creates a request to invoke DescribeEncryptionKey API
func CreateDescribeEncryptionKeyRequest() (request *DescribeEncryptionKeyRequest) {
	request = &DescribeEncryptionKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeEncryptionKey", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEncryptionKeyResponse creates a response to parse from DescribeEncryptionKey response
func CreateDescribeEncryptionKeyResponse() (response *DescribeEncryptionKeyResponse) {
	response = &DescribeEncryptionKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
