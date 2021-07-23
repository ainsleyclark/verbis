// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/ainsleyclark/verbis/api/domain"
	mock "github.com/stretchr/testify/mock"

	stow "github.com/graymeta/stow"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Bucket provides a mock function with given fields: provider, bucket
func (_m *Service) Bucket(provider domain.StorageProvider, bucket string) (stow.Container, error) {
	ret := _m.Called(provider, bucket)

	var r0 stow.Container
	if rf, ok := ret.Get(0).(func(domain.StorageProvider, string) stow.Container); ok {
		r0 = rf(provider, bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stow.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.StorageProvider, string) error); ok {
		r1 = rf(provider, bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BucketByFile provides a mock function with given fields: file
func (_m *Service) BucketByFile(file domain.File) (stow.Container, error) {
	ret := _m.Called(file)

	var r0 stow.Container
	if rf, ok := ret.Get(0).(func(domain.File) stow.Container); ok {
		r0 = rf(file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stow.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.File) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Config provides a mock function with given fields:
func (_m *Service) Config() (domain.StorageProvider, string, error) {
	ret := _m.Called()

	var r0 domain.StorageProvider
	if rf, ok := ret.Get(0).(func() domain.StorageProvider); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(domain.StorageProvider)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Provider provides a mock function with given fields: provider
func (_m *Service) Provider(provider domain.StorageProvider) (stow.Location, error) {
	ret := _m.Called(provider)

	var r0 stow.Location
	if rf, ok := ret.Get(0).(func(domain.StorageProvider) stow.Location); ok {
		r0 = rf(provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stow.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.StorageProvider) error); ok {
		r1 = rf(provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
