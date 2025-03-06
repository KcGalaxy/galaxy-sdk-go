// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ksmsiface provides an interface to enable mocking the ksms service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ksmsiface

import (
	"github.com/gexue/galaxy-sdk-go/service/ksms"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// KsmsAPI provides an interface to enable mocking the
// ksms.Ksms service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // ksms.
//    func myFunc(svc ksmsiface.KsmsAPI) bool {
//        // Make svc.BatchSendSms request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ksms.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKsmsClient struct {
//        ksmsiface.KsmsAPI
//    }
//    func (m *mockKsmsClient) BatchSendSms(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKsmsClient{}
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
type KsmsAPI interface {
	BatchSendSms(*map[string]interface{}) (*map[string]interface{}, error)
	BatchSendSmsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BatchSendSmsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SendSms(*map[string]interface{}) (*map[string]interface{}, error)
	SendSmsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SendSmsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ KsmsAPI = (*ksms.Ksms)(nil)
