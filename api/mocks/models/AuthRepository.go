// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: email, password
func (_m *AuthRepository) Authenticate(email string, password string) (domain.User, error) {
	ret := _m.Called(email, password)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string, string) domain.User); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanPasswordResets provides a mock function with given fields:
func (_m *AuthRepository) CleanPasswordResets() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logout provides a mock function with given fields: token
func (_m *AuthRepository) Logout(token string) (int, error) {
	ret := _m.Called(token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: token, password
func (_m *AuthRepository) ResetPassword(token string, password string) error {
	ret := _m.Called(token, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(token, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendResetPassword provides a mock function with given fields: email
func (_m *AuthRepository) SendResetPassword(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyEmail provides a mock function with given fields: md5String
func (_m *AuthRepository) VerifyEmail(md5String string) error {
	ret := _m.Called(md5String)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(md5String)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyPasswordToken provides a mock function with given fields: token
func (_m *AuthRepository) VerifyPasswordToken(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
