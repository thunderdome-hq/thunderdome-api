package api

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Type string

func Error(code codes.Code, errorType Type, cause string, args ...interface{}) error {
	text := fmt.Sprintf(cause, args...)
	return status.Errorf(code, fmt.Sprintf("%s failed –– %s.", errorType, text))
}
