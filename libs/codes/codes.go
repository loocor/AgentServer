package codes

import (
	"google.golang.org/grpc/codes"
)

const (
	// Common
	Unknown         codes.Code = 1001
	InvalidArgument codes.Code = 1002

	// Reg
	PhoneAlreadyExists    codes.Code = 2001
	IDNumberAlreadyExists codes.Code = 2002

	// Login
	PasswordWrong codes.Code = 2011
	UserNotFound  codes.Code = 2012
)
