package ddosbgp

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

// InstanceSpec is a nested struct in ddosbgp response
type InstanceSpec struct {
	IsFullDefenseMode             int        `json:"IsFullDefenseMode" xml:"IsFullDefenseMode"`
	Region                        string     `json:"Region" xml:"Region"`
	AvailableDefenseTimes         int        `json:"AvailableDefenseTimes" xml:"AvailableDefenseTimes"`
	TotalDefenseTimes             int        `json:"TotalDefenseTimes" xml:"TotalDefenseTimes"`
	AvailableDeleteBlackholeCount string     `json:"AvailableDeleteBlackholeCount" xml:"AvailableDeleteBlackholeCount"`
	InstanceId                    string     `json:"InstanceId" xml:"InstanceId"`
	PackConfig                    PackConfig `json:"PackConfig" xml:"PackConfig"`
}
