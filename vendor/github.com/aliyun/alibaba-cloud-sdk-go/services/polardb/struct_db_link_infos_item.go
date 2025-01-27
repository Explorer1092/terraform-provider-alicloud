package polardb

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

// DBLinkInfosItem is a nested struct in polardb response
type DBLinkInfosItem struct {
	DBInstanceName       string `json:"DBInstanceName" xml:"DBInstanceName"`
	DBLinkName           string `json:"DBLinkName" xml:"DBLinkName"`
	SourceDBName         string `json:"SourceDBName" xml:"SourceDBName"`
	TargetDBName         string `json:"TargetDBName" xml:"TargetDBName"`
	TargetDBInstanceName string `json:"TargetDBInstanceName" xml:"TargetDBInstanceName"`
	TargetAccount        string `json:"TargetAccount" xml:"TargetAccount"`
}
