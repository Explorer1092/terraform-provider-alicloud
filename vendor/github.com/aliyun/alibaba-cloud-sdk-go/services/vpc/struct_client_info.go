package vpc

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

// ClientInfo is a nested struct in vpc response
type ClientInfo struct {
	Status        string `json:"Status" xml:"Status"`
	PrivateIp     string `json:"PrivateIp" xml:"PrivateIp"`
	SendBytes     int64  `json:"SendBytes" xml:"SendBytes"`
	ConnectedTime int64  `json:"ConnectedTime" xml:"ConnectedTime"`
	CommonName    string `json:"CommonName" xml:"CommonName"`
	Ip            string `json:"Ip" xml:"Ip"`
	ReceiveBytes  int64  `json:"ReceiveBytes" xml:"ReceiveBytes"`
	Port          string `json:"Port" xml:"Port"`
}
