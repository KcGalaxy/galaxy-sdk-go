// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kcsv1iface provides an interface to enable mocking the kcsv1 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kcsv1iface

import (
	"github.com/KcGalaxy/galaxy-sdk-go/service/kcsv1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// Kcsv1API provides an interface to enable mocking the
// kcsv1.Kcsv1 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // kcsv1.
//    func myFunc(svc kcsv1iface.Kcsv1API) bool {
//        // Make svc.CreateCacheCluster request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kcsv1.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKcsv1Client struct {
//        kcsv1iface.Kcsv1API
//    }
//    func (m *mockKcsv1Client) CreateCacheCluster(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKcsv1Client{}
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
type Kcsv1API interface {
	CreateCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCacheParameterGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCacheParameterGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCacheParameterGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCacheSecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCacheSecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCacheSecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheParameterGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheParameterGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheParameterGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheSecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheSecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheSecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheSecurityGroupRule(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheSecurityGroupRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheSecurityGroupRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCacheSecurityRule(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCacheSecurityRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCacheSecurityRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheClusters(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheClustersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheClustersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheDefaultParameters(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheDefaultParametersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheDefaultParametersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheParameterGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheParameterGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheParameterGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheParameterGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheParameterGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheParameterGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheParameters(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheParametersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheParametersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheSecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheSecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheSecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheSecurityGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheSecurityGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheSecurityGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCacheSecurityRules(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCacheSecurityRulesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCacheSecurityRulesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSnapshots(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSnapshotsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSnapshotsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DownloadSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	DownloadSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DownloadSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ExportSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	ExportSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ExportSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	FlushCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	FlushCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	FlushCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCacheParameterGroup(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCacheParameterGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCacheParameterGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCacheSecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCacheSecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCacheSecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenameCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	RenameCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenameCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenameSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	RenameSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenameSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResizeCacheCluster(*map[string]interface{}) (*map[string]interface{}, error)
	ResizeCacheClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResizeCacheClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestoreSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	RestoreSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestoreSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetCacheParameterGroup(*map[string]interface{}) (*map[string]interface{}, error)
	SetCacheParameterGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetCacheParameterGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetCacheParameters(*map[string]interface{}) (*map[string]interface{}, error)
	SetCacheParametersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetCacheParametersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetCacheSecurityGroup(*map[string]interface{}) (*map[string]interface{}, error)
	SetCacheSecurityGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetCacheSecurityGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetCacheSecurityRules(*map[string]interface{}) (*map[string]interface{}, error)
	SetCacheSecurityRulesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetCacheSecurityRulesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetTimingSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	SetTimingSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetTimingSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdatePassword(*map[string]interface{}) (*map[string]interface{}, error)
	UpdatePasswordWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdatePasswordRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ Kcsv1API = (*kcsv1.Kcsv1)(nil)
