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

package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CPFSDescribeFileSystems invokes the nas.CPFSDescribeFileSystems API synchronously
// api document: https://help.aliyun.com/api/nas/cpfsdescribefilesystems.html
func (client *Client) CPFSDescribeFileSystems(request *CPFSDescribeFileSystemsRequest) (response *CPFSDescribeFileSystemsResponse, err error) {
	response = CreateCPFSDescribeFileSystemsResponse()
	err = client.DoAction(request, response)
	return
}

// CPFSDescribeFileSystemsWithChan invokes the nas.CPFSDescribeFileSystems API asynchronously
// api document: https://help.aliyun.com/api/nas/cpfsdescribefilesystems.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CPFSDescribeFileSystemsWithChan(request *CPFSDescribeFileSystemsRequest) (<-chan *CPFSDescribeFileSystemsResponse, <-chan error) {
	responseChan := make(chan *CPFSDescribeFileSystemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CPFSDescribeFileSystems(request)
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

// CPFSDescribeFileSystemsWithCallback invokes the nas.CPFSDescribeFileSystems API asynchronously
// api document: https://help.aliyun.com/api/nas/cpfsdescribefilesystems.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CPFSDescribeFileSystemsWithCallback(request *CPFSDescribeFileSystemsRequest, callback func(response *CPFSDescribeFileSystemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CPFSDescribeFileSystemsResponse
		var err error
		defer close(result)
		response, err = client.CPFSDescribeFileSystems(request)
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

// CPFSDescribeFileSystemsRequest is the request struct for api CPFSDescribeFileSystems
type CPFSDescribeFileSystemsRequest struct {
	*requests.RpcRequest
	FsId       string           `position:"Query" name:"FsId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// CPFSDescribeFileSystemsResponse is the response struct for api CPFSDescribeFileSystems
type CPFSDescribeFileSystemsResponse struct {
	*responses.BaseResponse
	RequestId   string                             `json:"RequestId" xml:"RequestId"`
	TotalCount  int                                `json:"TotalCount" xml:"TotalCount"`
	PageSize    int                                `json:"PageSize" xml:"PageSize"`
	PageNumber  int                                `json:"PageNumber" xml:"PageNumber"`
	FileSystems CPFSDescribeFileSystemsFileSystems `json:"FileSystems" xml:"FileSystems"`
}

type CPFSDescribeFileSystemsFileSystems struct {
	FileSystem []CPFSDescribeFileSystemsFileSystem `json:"FileSystem" xml:"FileSystem"`
}

type CPFSDescribeFileSystemsFileSystem struct {
	FsId         string `json:"fsId" xml:"fsId"`
	FsDesc       string `json:"fsDesc" xml:"fsDesc"`
	FsSpec       string `json:"fsSpec" xml:"fsSpec"`
	Bandwidth    int64  `json:"bandwidth" xml:"bandwidth"`
	Capacity     int64  `json:"capacity" xml:"capacity"`
	MeteredSize  int64  `json:"meteredSize" xml:"meteredSize"`
	CreateTime   string `json:"createTime" xml:"createTime"`
	NetworkType  string `json:"networkType" xml:"networkType"`
	MountTargets string `json:"mountTargets" xml:"mountTargets"`
	RegionId     string `json:"regionId" xml:"regionId"`
	ZoneId       string `json:"zoneId" xml:"zoneId"`
	VpcId        string `json:"vpcId" xml:"vpcId"`
	VSwitchId    string `json:"vSwitchId" xml:"vSwitchId"`
	Status       string `json:"status" xml:"status"`
}

// CreateCPFSDescribeFileSystemsRequest creates a request to invoke CPFSDescribeFileSystems API
func CreateCPFSDescribeFileSystemsRequest() (request *CPFSDescribeFileSystemsRequest) {
	request = &CPFSDescribeFileSystemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CPFSDescribeFileSystems", "nas", "openAPI")
	return
}

// CreateCPFSDescribeFileSystemsResponse creates a response to parse from CPFSDescribeFileSystems response
func CreateCPFSDescribeFileSystemsResponse() (response *CPFSDescribeFileSystemsResponse) {
	response = &CPFSDescribeFileSystemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
