package api

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	AuthError   = "authentication"
	StatusError = "status"
	JoinError   = "join"
	LeaveError  = "leave"
	AcceptError = "accept"
	VerifyError = "verify"
	ListError   = "list"
	ClaimError  = "claim"
	DropError   = "drop"
	CLIError    = "CLI"
)

func Error(code codes.Code, error string, cause string, args ...interface{}) error {
	text := fmt.Sprintf(cause, args)
	return status.Errorf(code, fmt.Sprintf("%s failed –– %s.", error, text))
}
