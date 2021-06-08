package ez_grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JSONStatus struct {
	Code    codes.Code    `json:"code"`
	Message string        `json:"message,omitempty"`
	Details []interface{} `json:"details"`
}

func ToErrorJSONStatus(err error) interface{} {
	s := status.Convert(err)
	if s == nil {
		return nil
	}
	return JSONStatus{
		Code:    s.Code(),
		Message: s.Message(),
		Details: s.Details(),
	}
}
