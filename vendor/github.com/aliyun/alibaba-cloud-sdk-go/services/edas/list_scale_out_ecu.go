package edas

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

// ListScaleOutEcu invokes the edas.ListScaleOutEcu API synchronously
func (client *Client) ListScaleOutEcu(request *ListScaleOutEcuRequest) (response *ListScaleOutEcuResponse, err error) {
	response = CreateListScaleOutEcuResponse()
	err = client.DoAction(request, response)
	return
}

// ListScaleOutEcuWithChan invokes the edas.ListScaleOutEcu API asynchronously
func (client *Client) ListScaleOutEcuWithChan(request *ListScaleOutEcuRequest) (<-chan *ListScaleOutEcuResponse, <-chan error) {
	responseChan := make(chan *ListScaleOutEcuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListScaleOutEcu(request)
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

// ListScaleOutEcuWithCallback invokes the edas.ListScaleOutEcu API asynchronously
func (client *Client) ListScaleOutEcuWithCallback(request *ListScaleOutEcuRequest, callback func(response *ListScaleOutEcuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListScaleOutEcuResponse
		var err error
		defer close(result)
		response, err = client.ListScaleOutEcu(request)
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

// ListScaleOutEcuRequest is the request struct for api ListScaleOutEcu
type ListScaleOutEcuRequest struct {
	*requests.RoaRequest
	Mem             string `position:"Query" name:"Mem"`
	LogicalRegionId string `position:"Query" name:"LogicalRegionId"`
	AppId           string `position:"Query" name:"AppId"`
	GroupId         string `position:"Query" name:"GroupId"`
	InstanceNum     string `position:"Query" name:"InstanceNum"`
	Cpu             string `position:"Query" name:"Cpu"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

// ListScaleOutEcuResponse is the response struct for api ListScaleOutEcu
type ListScaleOutEcuResponse struct {
	*responses.BaseResponse
	Code        int                          `json:"Code" xml:"Code"`
	Message     string                       `json:"Message" xml:"Message"`
	RequestId   string                       `json:"RequestId" xml:"RequestId"`
	EcuInfoList EcuInfoListInListScaleOutEcu `json:"EcuInfoList" xml:"EcuInfoList"`
}

// CreateListScaleOutEcuRequest creates a request to invoke ListScaleOutEcu API
func CreateListScaleOutEcuRequest() (request *ListScaleOutEcuRequest) {
	request = &ListScaleOutEcuRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListScaleOutEcu", "/pop/v5/resource/scale_out_ecu_list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListScaleOutEcuResponse creates a response to parse from ListScaleOutEcu response
func CreateListScaleOutEcuResponse() (response *ListScaleOutEcuResponse) {
	response = &ListScaleOutEcuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
