// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package roles

import (
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	mocks "github.com/verbiscms/verbis/api/mocks/store/roles"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// EditorTestSuite defines the helper used for roles
// testing.
type EditorTestSuite struct {
	test.HandlerSuite
}

// TestRoles
//
// Assert testing has begun.
func TestRoles(t *testing.T) {
	suite.Run(t, &EditorTestSuite{
		HandlerSuite: test.NewHandlerSuite(),
	})
}

// Setup
//
// A helper to obtain a mock editor handler
// for testing.
func (t *EditorTestSuite) Setup(mf func(m *mocks.Repository)) *Roles {
	m := &mocks.Repository{}
	if mf != nil {
		mf(m)
	}
	d := &deps.Deps{
		Store: &store.Repository{},
	}
	return New(d)
}
