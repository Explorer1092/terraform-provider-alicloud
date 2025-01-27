package bssopenapi

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

// QuerySavingsPlansDiscount invokes the bssopenapi.QuerySavingsPlansDiscount API synchronously
func (client *Client) QuerySavingsPlansDiscount(request *QuerySavingsPlansDiscountRequest) (response *QuerySavingsPlansDiscountResponse, err error) {
	response = CreateQuerySavingsPlansDiscountResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySavingsPlansDiscountWithChan invokes the bssopenapi.QuerySavingsPlansDiscount API asynchronously
func (client *Client) QuerySavingsPlansDiscountWithChan(request *QuerySavingsPlansDiscountRequest) (<-chan *QuerySavingsPlansDiscountResponse, <-chan error) {
	responseChan := make(chan *QuerySavingsPlansDiscountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySavingsPlansDiscount(request)
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

// QuerySavingsPlansDiscountWithCallback invokes the bssopenapi.QuerySavingsPlansDiscount API asynchronously
func (client *Client) QuerySavingsPlansDiscountWithCallback(request *QuerySavingsPlansDiscountRequest, callback func(response *QuerySavingsPlansDiscountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySavingsPlansDiscountResponse
		var err error
		defer close(result)
		response, err = client.QuerySavingsPlansDiscount(request)
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

// QuerySavingsPlansDiscountRequest is the request struct for api QuerySavingsPlansDiscount
type QuerySavingsPlansDiscountRequest struct {
	*requests.RpcRequest
	CommodityCode string           `position:"Query" name:"CommodityCode"`
	Locale        string           `position:"Query" name:"Locale"`
	PageNum       requests.Integer `position:"Query" name:"PageNum"`
	Cycle         string           `position:"Query" name:"Cycle"`
	Spec          string           `position:"Query" name:"Spec"`
	ModuleCode    string           `position:"Query" name:"ModuleCode"`
	PayMode       string           `position:"Query" name:"PayMode"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	SpnType       string           `position:"Query" name:"SpnType"`
	Region        string           `position:"Query" name:"Region"`
}

// QuerySavingsPlansDiscountResponse is the response struct for api QuerySavingsPlansDiscount
type QuerySavingsPlansDiscountResponse struct {
	*responses.BaseResponse
	Message   string                          `json:"Message" xml:"Message"`
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Code      string                          `json:"Code" xml:"Code"`
	Success   bool                            `json:"Success" xml:"Success"`
	Data      DataInQuerySavingsPlansDiscount `json:"Data" xml:"Data"`
}

// CreateQuerySavingsPlansDiscountRequest creates a request to invoke QuerySavingsPlansDiscount API
func CreateQuerySavingsPlansDiscountRequest() (request *QuerySavingsPlansDiscountRequest) {
	request = &QuerySavingsPlansDiscountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QuerySavingsPlansDiscount", "bssopenapi", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQuerySavingsPlansDiscountResponse creates a response to parse from QuerySavingsPlansDiscount response
func CreateQuerySavingsPlansDiscountResponse() (response *QuerySavingsPlansDiscountResponse) {
	response = &QuerySavingsPlansDiscountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
