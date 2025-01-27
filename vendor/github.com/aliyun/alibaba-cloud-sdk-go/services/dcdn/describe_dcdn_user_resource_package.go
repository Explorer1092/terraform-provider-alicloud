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

// DescribeDcdnUserResourcePackage invokes the dcdn.DescribeDcdnUserResourcePackage API synchronously
func (client *Client) DescribeDcdnUserResourcePackage(request *DescribeDcdnUserResourcePackageRequest) (response *DescribeDcdnUserResourcePackageResponse, err error) {
	response = CreateDescribeDcdnUserResourcePackageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnUserResourcePackageWithChan invokes the dcdn.DescribeDcdnUserResourcePackage API asynchronously
func (client *Client) DescribeDcdnUserResourcePackageWithChan(request *DescribeDcdnUserResourcePackageRequest) (<-chan *DescribeDcdnUserResourcePackageResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnUserResourcePackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnUserResourcePackage(request)
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

// DescribeDcdnUserResourcePackageWithCallback invokes the dcdn.DescribeDcdnUserResourcePackage API asynchronously
func (client *Client) DescribeDcdnUserResourcePackageWithCallback(request *DescribeDcdnUserResourcePackageRequest, callback func(response *DescribeDcdnUserResourcePackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnUserResourcePackageResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnUserResourcePackage(request)
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

// DescribeDcdnUserResourcePackageRequest is the request struct for api DescribeDcdnUserResourcePackage
type DescribeDcdnUserResourcePackageRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Status        string           `position:"Query" name:"Status"`
}

// DescribeDcdnUserResourcePackageResponse is the response struct for api DescribeDcdnUserResourcePackage
type DescribeDcdnUserResourcePackageResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	ResourcePackageInfos ResourcePackageInfos `json:"ResourcePackageInfos" xml:"ResourcePackageInfos"`
}

// CreateDescribeDcdnUserResourcePackageRequest creates a request to invoke DescribeDcdnUserResourcePackage API
func CreateDescribeDcdnUserResourcePackageRequest() (request *DescribeDcdnUserResourcePackageRequest) {
	request = &DescribeDcdnUserResourcePackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserResourcePackage", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnUserResourcePackageResponse creates a response to parse from DescribeDcdnUserResourcePackage response
func CreateDescribeDcdnUserResourcePackageResponse() (response *DescribeDcdnUserResourcePackageResponse) {
	response = &DescribeDcdnUserResourcePackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
