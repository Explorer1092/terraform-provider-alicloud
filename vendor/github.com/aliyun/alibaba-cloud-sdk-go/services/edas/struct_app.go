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

// App is a nested struct in edas response
type App struct {
	RequestMem              int     `json:"RequestMem" xml:"RequestMem"`
	InstancesBeforeScaling  int     `json:"InstancesBeforeScaling" xml:"InstancesBeforeScaling"`
	DeployType              string  `json:"DeployType" xml:"DeployType"`
	ApplicationName         string  `json:"ApplicationName" xml:"ApplicationName"`
	ApplicationType         string  `json:"ApplicationType" xml:"ApplicationType"`
	Instances               int     `json:"Instances" xml:"Instances"`
	LimitMem                int     `json:"LimitMem" xml:"LimitMem"`
	Cmd                     string  `json:"Cmd" xml:"Cmd"`
	RegionId                string  `json:"RegionId" xml:"RegionId"`
	BuildpackId             int     `json:"BuildpackId" xml:"BuildpackId"`
	TomcatVersion           string  `json:"TomcatVersion" xml:"TomcatVersion"`
	CsClusterId             string  `json:"CsClusterId" xml:"CsClusterId"`
	RequestCpuM             int     `json:"RequestCpuM" xml:"RequestCpuM"`
	AppId                   string  `json:"AppId" xml:"AppId"`
	K8sNamespace            string  `json:"K8sNamespace" xml:"K8sNamespace"`
	EdasContainerVersion    string  `json:"EdasContainerVersion" xml:"EdasContainerVersion"`
	LimitCpuM               int     `json:"LimitCpuM" xml:"LimitCpuM"`
	ClusterId               string  `json:"ClusterId" xml:"ClusterId"`
	DevelopType             string  `json:"DevelopType" xml:"DevelopType"`
	SlbInfo                 string  `json:"SlbInfo" xml:"SlbInfo"`
	Annotations             string  `json:"Annotations" xml:"Annotations"`
	Labels                  string  `json:"Labels" xml:"Labels"`
	LimitEphemeralStorage   string  `json:"LimitEphemeralStorage" xml:"LimitEphemeralStorage"`
	RequestEphemeralStorage string  `json:"RequestEphemeralStorage" xml:"RequestEphemeralStorage"`
	CmdArgs                 CmdArgs `json:"CmdArgs" xml:"CmdArgs"`
	EnvList                 EnvList `json:"EnvList" xml:"EnvList"`
}
