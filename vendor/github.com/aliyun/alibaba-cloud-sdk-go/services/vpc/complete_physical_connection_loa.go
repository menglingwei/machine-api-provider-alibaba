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

// CompletePhysicalConnectionLOA invokes the vpc.CompletePhysicalConnectionLOA API synchronously
func (client *Client) CompletePhysicalConnectionLOA(request *CompletePhysicalConnectionLOARequest) (response *CompletePhysicalConnectionLOAResponse, err error) {
	response = CreateCompletePhysicalConnectionLOAResponse()
	err = client.DoAction(request, response)
	return
}

// CompletePhysicalConnectionLOAWithChan invokes the vpc.CompletePhysicalConnectionLOA API asynchronously
func (client *Client) CompletePhysicalConnectionLOAWithChan(request *CompletePhysicalConnectionLOARequest) (<-chan *CompletePhysicalConnectionLOAResponse, <-chan error) {
	responseChan := make(chan *CompletePhysicalConnectionLOAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompletePhysicalConnectionLOA(request)
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

// CompletePhysicalConnectionLOAWithCallback invokes the vpc.CompletePhysicalConnectionLOA API asynchronously
func (client *Client) CompletePhysicalConnectionLOAWithCallback(request *CompletePhysicalConnectionLOARequest, callback func(response *CompletePhysicalConnectionLOAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompletePhysicalConnectionLOAResponse
		var err error
		defer close(result)
		response, err = client.CompletePhysicalConnectionLOA(request)
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

// CompletePhysicalConnectionLOARequest is the request struct for api CompletePhysicalConnectionLOA
type CompletePhysicalConnectionLOARequest struct {
	*requests.RpcRequest
	LineCode             string           `position:"Query" name:"LineCode"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	LineLabel            string           `position:"Query" name:"LineLabel"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// CompletePhysicalConnectionLOAResponse is the response struct for api CompletePhysicalConnectionLOA
type CompletePhysicalConnectionLOAResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCompletePhysicalConnectionLOARequest creates a request to invoke CompletePhysicalConnectionLOA API
func CreateCompletePhysicalConnectionLOARequest() (request *CompletePhysicalConnectionLOARequest) {
	request = &CompletePhysicalConnectionLOARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CompletePhysicalConnectionLOA", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCompletePhysicalConnectionLOAResponse creates a response to parse from CompletePhysicalConnectionLOA response
func CreateCompletePhysicalConnectionLOAResponse() (response *CompletePhysicalConnectionLOAResponse) {
	response = &CompletePhysicalConnectionLOAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
