package errors

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LimitSurpassionError struct {
	Method string
}

func (err LimitSurpassionError) Error() error {
	return status.Errorf(codes.NotFound, fmt.Sprintf("%s Your request is declined due to server requests overflow.", err.Method))
}
