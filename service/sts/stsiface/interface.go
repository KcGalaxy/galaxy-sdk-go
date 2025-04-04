// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package stsiface provides an interface to enable mocking the sts service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package stsiface

import (
	"github.com/KcGalaxy/galaxy-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// StsAPI provides an interface to enable mocking the
// sts.Sts service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // sts.
//    func myFunc(svc stsiface.StsAPI) bool {
//        // Make svc.AssumeRole request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sts.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockStsClient struct {
//        stsiface.StsAPI
//    }
//    func (m *mockStsClient) AssumeRole(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockStsClient{}
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
type StsAPI interface {
	AssumeRole(*map[string]interface{}) (*map[string]interface{}, error)
	AssumeRoleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssumeRoleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ StsAPI = (*sts.Sts)(nil)
