package errorcodes

import (
	"errors"
)

type ErrorCode string

type ValidationError struct {
	statusCode ErrorCode
	error      error
}

func (e ValidationError) StatusCode() ErrorCode {
	return e.statusCode
}

func (e ValidationError) Error() string {
	return e.error.Error()
}

const (
	ExpirationDateOther         ErrorCode = "100"
	ExpirationDateInvalidFormat ErrorCode = "101"
	Expired                     ErrorCode = "102"

	CardNumberOther         ErrorCode = "200"
	CardNumberInvalidFormat ErrorCode = "201"
	CardNumberLuhnFailed    ErrorCode = "202"
	CardNumberInvalidIIN    ErrorCode = "203"
)

var (
	ErrExpirationDate              = errors.New("expiration date validation failed")
	ErrExpirationDateInvalidFormat = errors.New("invalid expiration date format")
	ErrExpired                     = errors.New("card is expired")

	ErrInvalidCardNumber       = errors.New("card number validation failed")
	ErrInvalidCardNumberFormat = errors.New("invalid card number format")
	ErrCardNumberLuhnFailed    = errors.New("card number failed Luhn algorithm validation")
	ErrInvalidCardNumberIIN    = errors.New("invalid card IIN")
)

func Extract(err error) (ErrorCode, bool) {
	return "", false
}
