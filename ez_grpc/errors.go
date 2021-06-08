package ez_grpc

import (
	"github.com/levanthanh-ptit/go-ez/internal/ez_grpc"
)

func MakeInvalidArgument(err error) error {
	return ez_grpc.MakeInvalidArgument(err)
}

func MakeAlreadyExists(err error) error {
	return ez_grpc.MakeAlreadyExists(err)
}

func MakeDeadlineExceeded(err error) error {
	return ez_grpc.MakeDeadlineExceeded(err)
}

func MakePermissionDenied(err error) error {
	return ez_grpc.MakePermissionDenied(err)
}

func MakeUnauthenticated(err error) error {
	return ez_grpc.MakeUnauthenticated(err)
}

func MakeUnavailable(err error) error {
	return ez_grpc.MakeUnavailable(err)
}

func MakeNotFound(err error) error {
	return ez_grpc.MakeNotFound(err)
}
