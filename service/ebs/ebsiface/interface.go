// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ebsiface provides an interface to enable mocking the ebs service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ebsiface

import (
	"github.com/KcGalaxy/galaxy-sdk-go/service/ebs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// EbsAPI provides an interface to enable mocking the
// ebs.Ebs service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // ebs.
//    func myFunc(svc ebsiface.EbsAPI) bool {
//        // Make svc.AttachVolume request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ebs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEbsClient struct {
//        ebsiface.EbsAPI
//    }
//    func (m *mockEbsClient) AttachVolume(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEbsClient{}
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
type EbsAPI interface {
	AttachVolume(*map[string]interface{}) (*map[string]interface{}, error)
	AttachVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVolume(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVolume(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeEbsInstances(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeEbsInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeEbsInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceVolumes(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceVolumesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceVolumesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVolumes(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVolumesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVolumesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachVolume(*map[string]interface{}) (*map[string]interface{}, error)
	DetachVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVolume(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RecoveryVolume(*map[string]interface{}) (*map[string]interface{}, error)
	RecoveryVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RecoveryVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResizeVolume(*map[string]interface{}) (*map[string]interface{}, error)
	ResizeVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResizeVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateVolumeProject(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateVolumeProjectWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateVolumeProjectRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ValidateAttachInstance(*map[string]interface{}) (*map[string]interface{}, error)
	ValidateAttachInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ValidateAttachInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ EbsAPI = (*ebs.Ebs)(nil)
