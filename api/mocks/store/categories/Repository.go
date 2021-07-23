// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/verbiscms/verbis/api/domain"
	categories "github.com/verbiscms/verbis/api/store/categories"

	mock "github.com/stretchr/testify/mock"

	params "github.com/verbiscms/verbis/api/common/params"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *Repository) Create(c domain.Category) (domain.Category, error) {
	ret := _m.Called(c)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(domain.Category) domain.Category); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Category) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: id
func (_m *Repository) Exists(id int) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ExistsByName provides a mock function with given fields: name
func (_m *Repository) ExistsByName(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ExistsBySlug provides a mock function with given fields: slug
func (_m *Repository) ExistsBySlug(slug string) bool {
	ret := _m.Called(slug)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Find provides a mock function with given fields: id
func (_m *Repository) Find(id int) (domain.Category, error) {
	ret := _m.Called(id)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(int) domain.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *Repository) FindByName(name string) (domain.Category, error) {
	ret := _m.Called(name)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(string) domain.Category); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPost provides a mock function with given fields: id
func (_m *Repository) FindByPost(id int) (domain.Category, error) {
	ret := _m.Called(id)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(int) domain.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBySlug provides a mock function with given fields: slug
func (_m *Repository) FindBySlug(slug string) (domain.Category, error) {
	ret := _m.Called(slug)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(string) domain.Category); ok {
		r0 = rf(slug)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindParent provides a mock function with given fields: id
func (_m *Repository) FindParent(id int) (domain.Category, error) {
	ret := _m.Called(id)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(int) domain.Category); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: meta, cfg
func (_m *Repository) List(meta params.Params, cfg categories.ListConfig) (domain.Categories, int, error) {
	ret := _m.Called(meta, cfg)

	var r0 domain.Categories
	if rf, ok := ret.Get(0).(func(params.Params, categories.ListConfig) domain.Categories); ok {
		r0 = rf(meta, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Categories)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(params.Params, categories.ListConfig) int); ok {
		r1 = rf(meta, cfg)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(params.Params, categories.ListConfig) error); ok {
		r2 = rf(meta, cfg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: c
func (_m *Repository) Update(c domain.Category) (domain.Category, error) {
	ret := _m.Called(c)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(domain.Category) domain.Category); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Category) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
