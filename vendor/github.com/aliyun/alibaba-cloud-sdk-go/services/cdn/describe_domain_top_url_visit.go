package cdn

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

// DescribeDomainTopUrlVisit invokes the cdn.DescribeDomainTopUrlVisit API synchronously
func (client *Client) DescribeDomainTopUrlVisit(request *DescribeDomainTopUrlVisitRequest) (response *DescribeDomainTopUrlVisitResponse, err error) {
	response = CreateDescribeDomainTopUrlVisitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainTopUrlVisitWithChan invokes the cdn.DescribeDomainTopUrlVisit API asynchronously
func (client *Client) DescribeDomainTopUrlVisitWithChan(request *DescribeDomainTopUrlVisitRequest) (<-chan *DescribeDomainTopUrlVisitResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainTopUrlVisitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainTopUrlVisit(request)
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

// DescribeDomainTopUrlVisitWithCallback invokes the cdn.DescribeDomainTopUrlVisit API asynchronously
func (client *Client) DescribeDomainTopUrlVisitWithCallback(request *DescribeDomainTopUrlVisitRequest, callback func(response *DescribeDomainTopUrlVisitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainTopUrlVisitResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainTopUrlVisit(request)
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

// DescribeDomainTopUrlVisitRequest is the request struct for api DescribeDomainTopUrlVisit
type DescribeDomainTopUrlVisitRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	SortBy     string `position:"Query" name:"SortBy"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainTopUrlVisitResponse is the response struct for api DescribeDomainTopUrlVisit
type DescribeDomainTopUrlVisitResponse struct {
	*responses.BaseResponse
	StartTime  string                                `json:"StartTime" xml:"StartTime"`
	RequestId  string                                `json:"RequestId" xml:"RequestId"`
	DomainName string                                `json:"DomainName" xml:"DomainName"`
	AllUrlList AllUrlListInDescribeDomainTopUrlVisit `json:"AllUrlList" xml:"AllUrlList"`
	Url200List Url200ListInDescribeDomainTopUrlVisit `json:"Url200List" xml:"Url200List"`
	Url300List Url300ListInDescribeDomainTopUrlVisit `json:"Url300List" xml:"Url300List"`
	Url400List Url400ListInDescribeDomainTopUrlVisit `json:"Url400List" xml:"Url400List"`
	Url500List Url500ListInDescribeDomainTopUrlVisit `json:"Url500List" xml:"Url500List"`
}

// CreateDescribeDomainTopUrlVisitRequest creates a request to invoke DescribeDomainTopUrlVisit API
func CreateDescribeDomainTopUrlVisitRequest() (request *DescribeDomainTopUrlVisitRequest) {
	request = &DescribeDomainTopUrlVisitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainTopUrlVisit", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainTopUrlVisitResponse creates a response to parse from DescribeDomainTopUrlVisit response
func CreateDescribeDomainTopUrlVisitResponse() (response *DescribeDomainTopUrlVisitResponse) {
	response = &DescribeDomainTopUrlVisitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
