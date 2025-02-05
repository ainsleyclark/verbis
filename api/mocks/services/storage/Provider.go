// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	domain "github.com/verbiscms/verbis/api/domain"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/verbiscms/verbis/api/services/storage"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// Connect provides a mock function with given fields: info
func (_m *Provider) Connect(info domain.StorageConfig) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.StorageConfig) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBucket provides a mock function with given fields: provider, name
func (_m *Provider) CreateBucket(provider domain.StorageProvider, name string) (domain.Bucket, error) {
	ret := _m.Called(provider, name)

	var r0 domain.Bucket
	if rf, ok := ret.Get(0).(func(domain.StorageProvider, string) domain.Bucket); ok {
		r0 = rf(provider, name)
	} else {
		r0 = ret.Get(0).(domain.Bucket)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.StorageProvider, string) error); ok {
		r1 = rf(provider, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Provider) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBucket provides a mock function with given fields: provider, name
func (_m *Provider) DeleteBucket(provider domain.StorageProvider, name string) error {
	ret := _m.Called(provider, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.StorageProvider, string) error); ok {
		r0 = rf(provider, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields:
func (_m *Provider) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Download provides a mock function with given fields: w
func (_m *Provider) Download(w io.Writer) error {
	ret := _m.Called(w)

	var r0 error
	if rf, ok := ret.Get(0).(func(io.Writer) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: name
func (_m *Provider) Exists(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: url
func (_m *Provider) Find(url string) ([]byte, domain.File, error) {
	ret := _m.Called(url)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 domain.File
	if rf, ok := ret.Get(1).(func(string) domain.File); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Get(1).(domain.File)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(url)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Info provides a mock function with given fields: ctx
func (_m *Provider) Info(ctx context.Context) storage.Configuration {
	ret := _m.Called(ctx)

	var r0 storage.Configuration
	if rf, ok := ret.Get(0).(func(context.Context) storage.Configuration); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(storage.Configuration)
	}

	return r0
}

// ListBuckets provides a mock function with given fields: provider
func (_m *Provider) ListBuckets(provider domain.StorageProvider) (domain.Buckets, error) {
	ret := _m.Called(provider)

	var r0 domain.Buckets
	if rf, ok := ret.Get(0).(func(domain.StorageProvider) domain.Buckets); ok {
		r0 = rf(provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Buckets)
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

// Migrate provides a mock function with given fields: ctx, server, delete
func (_m *Provider) Migrate(ctx context.Context, server bool, delete bool) (int, error) {
	ret := _m.Called(ctx, server, delete)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, bool, bool) int); ok {
		r0 = rf(ctx, server, delete)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool, bool) error); ok {
		r1 = rf(ctx, server, delete)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: upload
func (_m *Provider) Upload(upload domain.Upload) (domain.File, error) {
	ret := _m.Called(upload)

	var r0 domain.File
	if rf, ok := ret.Get(0).(func(domain.Upload) domain.File); ok {
		r0 = rf(upload)
	} else {
		r0 = ret.Get(0).(domain.File)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Upload) error); ok {
		r1 = rf(upload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
