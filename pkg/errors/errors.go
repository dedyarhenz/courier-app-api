package errors

import "errors"

var ErrInvalidRequest = errors.New("invalid request")
var ErrLoginFailed = errors.New("wrong email or password")
var ErrEmailAlready = errors.New("email already exist")

var ErrUserNotFound = errors.New("user not found")
var ErrSizeNotFound = errors.New("size not found")
var ErrCategoryNotFound = errors.New("category not found")
var ErrAddOnNotFound = errors.New("add on not found")
var ErrAddressNotFound = errors.New("address not found")

var ErrMinTopUp = errors.New("topup min Rp.10.000")
var ErrMaxTopUp = errors.New("topup max Rp.10.000.000")

var ErrMinAddOns = errors.New("min select 1 add on")
var ErrAddOnInvalid = errors.New("add on id invalid")
