package drds

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

// Data is a nested struct in drds response
type Data struct {
	Version            int64                      `json:"Version" xml:"Version"`
	SpecTypeName       string                     `json:"SpecTypeName" xml:"SpecTypeName"`
	OrderId            int64                      `json:"OrderId" xml:"OrderId"`
	CreateTime         int64                      `json:"CreateTime" xml:"CreateTime"`
	DbName             string                     `json:"DbName" xml:"DbName"`
	DrdsInstanceId     string                     `json:"DrdsInstanceId" xml:"DrdsInstanceId"`
	VpcCloudInstanceId string                     `json:"VpcCloudInstanceId" xml:"VpcCloudInstanceId"`
	RegionId           string                     `json:"RegionId" xml:"RegionId"`
	ZoneId             string                     `json:"ZoneId" xml:"ZoneId"`
	AccountName        string                     `json:"AccountName" xml:"AccountName"`
	SpecTypeId         string                     `json:"SpecTypeId" xml:"SpecTypeId"`
	Status             string                     `json:"Status" xml:"Status"`
	NetworkType        string                     `json:"NetworkType" xml:"NetworkType"`
	Description        string                     `json:"Description" xml:"Description"`
	Specification      string                     `json:"Specification" xml:"Specification"`
	Type               string                     `json:"Type" xml:"Type"`
	DrdsInstanceIdList DrdsInstanceIdList         `json:"DrdsInstanceIdList" xml:"DrdsInstanceIdList"`
	IpWhiteList        IpWhiteList                `json:"IpWhiteList" xml:"IpWhiteList"`
	Vips               VipsInDescribeDrdsInstance `json:"Vips" xml:"Vips"`
}
