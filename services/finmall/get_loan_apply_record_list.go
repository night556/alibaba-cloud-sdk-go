package finmall

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

// GetLoanApplyRecordList invokes the finmall.GetLoanApplyRecordList API synchronously
// api document: https://help.aliyun.com/api/finmall/getloanapplyrecordlist.html
func (client *Client) GetLoanApplyRecordList(request *GetLoanApplyRecordListRequest) (response *GetLoanApplyRecordListResponse, err error) {
	response = CreateGetLoanApplyRecordListResponse()
	err = client.DoAction(request, response)
	return
}

// GetLoanApplyRecordListWithChan invokes the finmall.GetLoanApplyRecordList API asynchronously
// api document: https://help.aliyun.com/api/finmall/getloanapplyrecordlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLoanApplyRecordListWithChan(request *GetLoanApplyRecordListRequest) (<-chan *GetLoanApplyRecordListResponse, <-chan error) {
	responseChan := make(chan *GetLoanApplyRecordListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLoanApplyRecordList(request)
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

// GetLoanApplyRecordListWithCallback invokes the finmall.GetLoanApplyRecordList API asynchronously
// api document: https://help.aliyun.com/api/finmall/getloanapplyrecordlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLoanApplyRecordListWithCallback(request *GetLoanApplyRecordListRequest, callback func(response *GetLoanApplyRecordListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLoanApplyRecordListResponse
		var err error
		defer close(result)
		response, err = client.GetLoanApplyRecordList(request)
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

// GetLoanApplyRecordListRequest is the request struct for api GetLoanApplyRecordList
type GetLoanApplyRecordListRequest struct {
	*requests.RpcRequest
	CreditId string `position:"Query" name:"CreditId"`
	UserId   string `position:"Query" name:"UserId"`
}

// GetLoanApplyRecordListResponse is the response struct for api GetLoanApplyRecordList
type GetLoanApplyRecordListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetLoanApplyRecordListRequest creates a request to invoke GetLoanApplyRecordList API
func CreateGetLoanApplyRecordListRequest() (request *GetLoanApplyRecordListRequest) {
	request = &GetLoanApplyRecordListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "GetLoanApplyRecordList", "finmall", "openAPI")
	return
}

// CreateGetLoanApplyRecordListResponse creates a response to parse from GetLoanApplyRecordList response
func CreateGetLoanApplyRecordListResponse() (response *GetLoanApplyRecordListResponse) {
	response = &GetLoanApplyRecordListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
