// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/verbiscms/verbis/api/domain"
	environment "github.com/verbiscms/verbis/api/environment"

	mock "github.com/stretchr/testify/mock"

	stow "github.com/graymeta/stow"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// Dial provides a mock function with given fields: env
func (_m *Provider) Dial(env *environment.Env) (stow.Location, error) {
	ret := _m.Called(env)

	var r0 stow.Location
	if rf, ok := ret.Get(0).(func(*environment.Env) stow.Location); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stow.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*environment.Env) error); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Info provides a mock function with given fields: env
func (_m *Provider) Info(env *environment.Env) domain.StorageProviderInfo {
	ret := _m.Called(env)

	var r0 domain.StorageProviderInfo
	if rf, ok := ret.Get(0).(func(*environment.Env) domain.StorageProviderInfo); ok {
		r0 = rf(env)
	} else {
		r0 = ret.Get(0).(domain.StorageProviderInfo)
	}

	return r0
}
