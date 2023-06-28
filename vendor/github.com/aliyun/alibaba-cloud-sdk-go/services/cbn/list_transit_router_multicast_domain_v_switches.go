package cbn

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

// ListTransitRouterMulticastDomainVSwitches invokes the cbn.ListTransitRouterMulticastDomainVSwitches API synchronously
func (client *Client) ListTransitRouterMulticastDomainVSwitches(request *ListTransitRouterMulticastDomainVSwitchesRequest) (response *ListTransitRouterMulticastDomainVSwitchesResponse, err error) {
	response = CreateListTransitRouterMulticastDomainVSwitchesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterMulticastDomainVSwitchesWithChan invokes the cbn.ListTransitRouterMulticastDomainVSwitches API asynchronously
func (client *Client) ListTransitRouterMulticastDomainVSwitchesWithChan(request *ListTransitRouterMulticastDomainVSwitchesRequest) (<-chan *ListTransitRouterMulticastDomainVSwitchesResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterMulticastDomainVSwitchesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterMulticastDomainVSwitches(request)
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

// ListTransitRouterMulticastDomainVSwitchesWithCallback invokes the cbn.ListTransitRouterMulticastDomainVSwitches API asynchronously
func (client *Client) ListTransitRouterMulticastDomainVSwitchesWithCallback(request *ListTransitRouterMulticastDomainVSwitchesRequest, callback func(response *ListTransitRouterMulticastDomainVSwitchesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterMulticastDomainVSwitchesResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterMulticastDomainVSwitches(request)
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

// ListTransitRouterMulticastDomainVSwitchesRequest is the request struct for api ListTransitRouterMulticastDomainVSwitches
type ListTransitRouterMulticastDomainVSwitchesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	VSwitchIds           *[]string        `position:"Query" name:"VSwitchIds"  type:"Repeated"`
	NextToken            string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
	VpcId                string           `position:"Query" name:"VpcId"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
}

// ListTransitRouterMulticastDomainVSwitchesResponse is the response struct for api ListTransitRouterMulticastDomainVSwitches
type ListTransitRouterMulticastDomainVSwitchesResponse struct {
	*responses.BaseResponse
	NextToken  string   `json:"NextToken" xml:"NextToken"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	MaxResults int      `json:"MaxResults" xml:"MaxResults"`
	VSwitchIds []string `json:"VSwitchIds" xml:"VSwitchIds"`
}

// CreateListTransitRouterMulticastDomainVSwitchesRequest creates a request to invoke ListTransitRouterMulticastDomainVSwitches API
func CreateListTransitRouterMulticastDomainVSwitchesRequest() (request *ListTransitRouterMulticastDomainVSwitchesRequest) {
	request = &ListTransitRouterMulticastDomainVSwitchesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterMulticastDomainVSwitches", "", "")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterMulticastDomainVSwitchesResponse creates a response to parse from ListTransitRouterMulticastDomainVSwitches response
func CreateListTransitRouterMulticastDomainVSwitchesResponse() (response *ListTransitRouterMulticastDomainVSwitchesResponse) {
	response = &ListTransitRouterMulticastDomainVSwitchesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}