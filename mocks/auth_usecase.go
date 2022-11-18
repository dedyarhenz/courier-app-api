// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project-backend/dto"

	mock "github.com/stretchr/testify/mock"
)

// AuthUsecase is an autogenerated mock type for the AuthUsecase type
type AuthUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: request
func (_m *AuthUsecase) Login(request dto.UserLoginRequest) (string, error) {
	ret := _m.Called(request)

	var r0 string
	if rf, ok := ret.Get(0).(func(dto.UserLoginRequest) string); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.UserLoginRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: request
func (_m *AuthUsecase) Register(request dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
	ret := _m.Called(request)

	var r0 *dto.UserRegisterResponse
	if rf, ok := ret.Get(0).(func(dto.UserRegisterRequest) *dto.UserRegisterResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserRegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.UserRegisterRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthUsecase creates a new instance of AuthUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthUsecase(t mockConstructorTestingTNewAuthUsecase) *AuthUsecase {
	mock := &AuthUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
