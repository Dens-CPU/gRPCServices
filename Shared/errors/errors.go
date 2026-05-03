package sharederrors

import "errors"

// Errors from UserSerivce
var (
	ExpiredToken      = errors.New("the token has expired")
	UserBlocked       = errors.New("user is blocked")
	ReAutentification = errors.New("re-autentification required")
)

// Errors from OrderService
var ()

//Errors from SpotSerivce
