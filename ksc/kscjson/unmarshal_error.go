package kscjson

import (
	"github.com/KcGalaxy/galaxy-sdk-go/ksc/kscquery"
	"github.com/aws/aws-sdk-go/aws/request"
)

type xmlErrorResponse struct {
	Error     Error  `json:"Error"`
	RequestID string `json:"RequestID"`
}

type Error struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

// UnmarshalErrorHandler is a name request handler to unmarshal request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "kscsdk.query.UnmarshalError", Fn: UnmarshalError}

// UnmarshalError unmarshals an error response for an AWS Query service.
func UnmarshalError(r *request.Request) {
	kscquery.UnmarshalError(r)
}
