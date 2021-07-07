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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyIPv6TranslatorEntry invokes the vpc.ModifyIPv6TranslatorEntry API synchronously
func (client *Client) ModifyIPv6TranslatorEntry(request *ModifyIPv6TranslatorEntryRequest) (response *ModifyIPv6TranslatorEntryResponse, err error) {
	response = CreateModifyIPv6TranslatorEntryResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIPv6TranslatorEntryWithChan invokes the vpc.ModifyIPv6TranslatorEntry API asynchronously
func (client *Client) ModifyIPv6TranslatorEntryWithChan(request *ModifyIPv6TranslatorEntryRequest) (<-chan *ModifyIPv6TranslatorEntryResponse, <-chan error) {
	responseChan := make(chan *ModifyIPv6TranslatorEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIPv6TranslatorEntry(request)
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

// ModifyIPv6TranslatorEntryWithCallback invokes the vpc.ModifyIPv6TranslatorEntry API asynchronously
func (client *Client) ModifyIPv6TranslatorEntryWithCallback(request *ModifyIPv6TranslatorEntryRequest, callback func(response *ModifyIPv6TranslatorEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIPv6TranslatorEntryResponse
		var err error
		defer close(result)
		response, err = client.ModifyIPv6TranslatorEntry(request)
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

// ModifyIPv6TranslatorEntryRequest is the request struct for api ModifyIPv6TranslatorEntry
type ModifyIPv6TranslatorEntryRequest struct {
	*requests.RpcRequest
	BackendIpv4Port       requests.Integer `position:"Query" name:"BackendIpv4Port"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EntryName             string           `position:"Query" name:"EntryName"`
	AclStatus             string           `position:"Query" name:"AclStatus"`
	EntryBandwidth        requests.Integer `position:"Query" name:"EntryBandwidth"`
	AclType               string           `position:"Query" name:"AclType"`
	AllocateIpv6Port      requests.Integer `position:"Query" name:"AllocateIpv6Port"`
	EntryDescription      string           `position:"Query" name:"EntryDescription"`
	BackendIpv4Addr       string           `position:"Query" name:"BackendIpv4Addr"`
	AclId                 string           `position:"Query" name:"AclId"`
	Ipv6TranslatorEntryId string           `position:"Query" name:"Ipv6TranslatorEntryId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	TransProtocol         string           `position:"Query" name:"TransProtocol"`
}

// ModifyIPv6TranslatorEntryResponse is the response struct for api ModifyIPv6TranslatorEntry
type ModifyIPv6TranslatorEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyIPv6TranslatorEntryRequest creates a request to invoke ModifyIPv6TranslatorEntry API
func CreateModifyIPv6TranslatorEntryRequest() (request *ModifyIPv6TranslatorEntryRequest) {
	request = &ModifyIPv6TranslatorEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyIPv6TranslatorEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyIPv6TranslatorEntryResponse creates a response to parse from ModifyIPv6TranslatorEntry response
func CreateModifyIPv6TranslatorEntryResponse() (response *ModifyIPv6TranslatorEntryResponse) {
	response = &ModifyIPv6TranslatorEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
