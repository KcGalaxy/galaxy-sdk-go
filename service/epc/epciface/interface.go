// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package epciface provides an interface to enable mocking the epc service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package epciface

import (
	"github.com/KcGalaxy/galaxy-sdk-go/service/epc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// EpcAPI provides an interface to enable mocking the
// epc.Epc service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // epc.
//    func myFunc(svc epciface.EpcAPI) bool {
//        // Make svc.AssociateCluster request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := epc.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEpcClient struct {
//        epciface.EpcAPI
//    }
//    func (m *mockEpcClient) AssociateCluster(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEpcClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type EpcAPI interface {
	AssociateCluster(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BuyAccessory(*map[string]interface{}) (*map[string]interface{}, error)
	BuyAccessoryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BuyAccessoryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAccessory(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAccessoryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAccessoryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCabinet(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCabinetWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCabinetRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateEpc(*map[string]interface{}) (*map[string]interface{}, error)
	CreateEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateImage(*map[string]interface{}) (*map[string]interface{}, error)
	CreateImageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateImageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateKey(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateProcess(*map[string]interface{}) (*map[string]interface{}, error)
	CreateProcessWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateProcessRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRemoteManagement(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRemoteManagementWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRemoteManagementRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAccessory(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAccessoryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAccessoryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteEpc(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteImage(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteImageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteImageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteKey(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteProcess(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteProcessWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteProcessRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRemoteManagement(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRemoteManagementWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRemoteManagementRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAccessorys(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAccessorysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAccessorysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCabinets(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCabinetsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCabinetsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeEpcDeviceAttributes(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeEpcDeviceAttributesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeEpcDeviceAttributesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeEpcManagements(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeEpcManagementsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeEpcManagementsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeEpcStocks(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeEpcStocksWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeEpcStocksRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeEpcs(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeEpcsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeEpcsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeImages(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeImagesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeImagesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInspections(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInspectionsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInspectionsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKeys(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKeysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKeysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribePhysicalMonitor(*map[string]interface{}) (*map[string]interface{}, error)
	DescribePhysicalMonitorWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribePhysicalMonitorRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeProcesses(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeProcessesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeProcessesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRemoteManagements(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRemoteManagementsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRemoteManagementsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpns(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateCluster(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetDynamicCode(*map[string]interface{}) (*map[string]interface{}, error)
	GetDynamicCodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetDynamicCodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ImportKey(*map[string]interface{}) (*map[string]interface{}, error)
	ImportKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ImportKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDns(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDnsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDnsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyEpc(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyHyperThreading(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyHyperThreadingWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyHyperThreadingRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyImage(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyImageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyImageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyKey(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNetworkInterfaceAttribute(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNetworkInterfaceAttributeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNetworkInterfaceAttributeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyRemoteManagement(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyRemoteManagementWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyRemoteManagementRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RebootEpc(*map[string]interface{}) (*map[string]interface{}, error)
	RebootEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RebootEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReinstallCustomerEpc(*map[string]interface{}) (*map[string]interface{}, error)
	ReinstallCustomerEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReinstallCustomerEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReinstallEpc(*map[string]interface{}) (*map[string]interface{}, error)
	ReinstallEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReinstallEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReplyProcess(*map[string]interface{}) (*map[string]interface{}, error)
	ReplyProcessWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReplyProcessRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResetPassword(*map[string]interface{}) (*map[string]interface{}, error)
	ResetPasswordWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResetPasswordRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartEpc(*map[string]interface{}) (*map[string]interface{}, error)
	StartEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopEpc(*map[string]interface{}) (*map[string]interface{}, error)
	StopEpcWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopEpcRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ EpcAPI = (*epc.Epc)(nil)
