package roles

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
	"testing"
)

// RolesTestSuite defines the helper used for role
// testing.
type RolesTestSuite struct {
	test.DBSuite
}

// TestRoles
//
// Assert testing has begun.
func TestRoles(t *testing.T) {
	suite.Run(t, &RolesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock roles database
// for testing.
func (t *RolesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

const (
	// The default role ID used for testing.
	roleID = "1"
)

var (
	// The default role used for testing.
	role = domain.Role{
		ID:          domain.BannedRoleID,
		Name:        "Banned",
		Description: "The user has been banned from the system.",
	}
	// The default roles used for testing.
	roles = domain.Roles{
		{
			ID:          domain.BannedRoleID,
			Name:        "Banned",
			Description: "The user has been banned from the system.",
		},
		{
			ID:          domain.ContributorRoleID,
			Name:        "Contributor",
			Description: "The user can create and edit their own draft posts, but they are unable to edit drafts of users or published posts.",
		},
	}
)
