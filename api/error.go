package api

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Type string

const (
	AuthError   Type = "authentication"
	StatusError Type = "status"
	JoinError   Type = "join"
	LeaveError  Type = "leave"
	AcceptError Type = "accept"
	VerifyError Type = "verify"
	ListError   Type = "list"
	ClaimError  Type = "claim"
	DropError   Type = "drop"
	CLIError    Type = "CLI"
)

func Error(code codes.Code, errorType Type, cause string, args ...interface{}) error {
	text := fmt.Sprintf(cause, args...)
	return status.Errorf(code, fmt.Sprintf("%s failed –– %s.", errorType, text))
}
