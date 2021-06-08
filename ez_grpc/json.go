package ez_grpc

import (
	"github.com/levanthanh-ptit/go-ez/internal/ez_grpc"
)

func ToErrorJSONStatus(err error) interface{} {
	return ez_grpc.ToErrorJSONStatus(err)
}
