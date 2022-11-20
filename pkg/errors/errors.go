package errors

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrLoginFailed = errors.New("wrong email or password")
var ErrInvalidRequest = errors.New("invalid request")
var ErrEmailAlready = errors.New("email already exist")

var ErrSizeNotFound = errors.New("size not found")
var ErrCategoryNotFound = errors.New("category not found")
