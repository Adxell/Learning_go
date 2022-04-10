package model 
import "errors"

var(
	ErrPersonCanNotBenil=errors.New("Person can not be null")
	ErrIDPersonDoesNotExit=errors.New("Person does not exit")
)