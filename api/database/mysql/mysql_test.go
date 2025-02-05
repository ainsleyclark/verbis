// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mysql

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	sm "github.com/hashicorp/go-version"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/database/builder"
	"github.com/verbiscms/verbis/api/database/internal"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/test"
	"github.com/verbiscms/verbis/api/version"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

// MySQLTestSuite defines the helper used for MySQL
// testing.
type MySQLTestSuite struct {
	test.DBSuite
	path string
}

// TestMySQL
//
// Assert testing has begun.
func TestMySQL(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return
	}
	suite.Run(t, &MySQLTestSuite{
		DBSuite: test.NewDBSuite(t),
		path:    path,
	})
}

// Setup
//
// A helper to obtain a mock MySQL database
// for testing.
func (t *MySQLTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *MySQL {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return &MySQL{
		driver: t.DB,
		env: &environment.Env{
			DbDriver:   "mysql",
			DbHost:     "127.0.0.1",
			DbPort:     "3306",
			DbDatabase: "verbis",
			DbUser:     "root",
			DbPassword: "password",
			DbSchema:   "",
		},
	}
}

func (t *MySQLTestSuite) SetupSuite() {
	db, m, err := sqlmock.New()
	t.NoError(err)
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	t.Mock = m
	t.DB = sqlxDB
}

func (t *MySQLTestSuite) TestMySQL_DB() {
	m := t.Setup(nil)
	t.Equal(t.DB, m.DB())
}

func (t *MySQLTestSuite) TestMySQL_Schema() {
	m := t.Setup(nil)
	t.Equal("", m.Schema())
}

func (t *MySQLTestSuite) TestMySQL_Builder() {
	m := t.Setup(nil)
	t.Equal(builder.New("mysql"), m.Builder())
}

func (t *MySQLTestSuite) TestMySQL_Close() {
	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Success": {
			func(m sqlmock.Sqlmock) {
				m.ExpectClose()
			},
			nil,
		},
		"Error": {
			func(m sqlmock.Sqlmock) {
				m.ExpectClose().
					WillReturnError(fmt.Errorf("error"))
			},
			"Error closing database",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := t.Setup(test.mock)
			err := m.Close()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(err, test.want)
		})
	}
}

type mockMigratorError struct{}

func (m *mockMigratorError) Migrate(ver *sm.Version, install bool) error {
	return fmt.Errorf("error")
}

type mockMigrator struct{}

func (m *mockMigrator) Migrate(ver *sm.Version, install bool) error {
	return nil
}

func (t *MySQLTestSuite) TestMySQL_Install() {
	tt := map[string]struct {
		mock internal.Migrator
		want interface{}
	}{
		"Success": {
			&mockMigrator{},
			nil,
		},
		"Error": {
			&mockMigratorError{},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := MySQL{migrator: test.mock}
			err := m.Install()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(err, test.want)
		})
	}
}

func (t *MySQLTestSuite) TestMySQL_Migrate() {
	tt := map[string]struct {
		mock internal.Migrator
		want interface{}
	}{
		"Success": {
			&mockMigrator{},
			nil,
		},
		"Error": {
			&mockMigratorError{},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := MySQL{migrator: test.mock}
			err := m.Migrate(version.SemVer)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(err, test.want)
		})
	}
}

func (t *MySQLTestSuite) TestMySQL_Tables() {
	var tbl []string

	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Success": {
			func(m sqlmock.Sqlmock) {
				for _, table := range internal.Tables {
					q := "SELECT EXISTS (SELECT * FROM `information_schema`.`tables` WHERE `table_schema` = 'verbis' AND `table_name` = '" + table + "' LIMIT 1)"
					rows := sqlmock.NewRows([]string{"id"}).
						AddRow(true)
					m.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)
				}
			},
			tbl,
		},
		"DB Error": {
			func(m sqlmock.Sqlmock) {
				for _, table := range internal.Tables {
					q := "SELECT EXISTS (SELECT * FROM `information_schema`.`tables` WHERE `table_schema` = 'verbis' AND `table_name` = '" + table + "' LIMIT 1)"
					m.ExpectQuery(regexp.QuoteMeta(q)).WillReturnError(fmt.Errorf("error"))
				}
			},
			"error",
		},
		"Failed Tables": {
			func(m sqlmock.Sqlmock) {
				for _, table := range internal.Tables {
					q := "SELECT EXISTS (SELECT * FROM `information_schema`.`tables` WHERE `table_schema` = 'verbis' AND `table_name` = '" + table + "' LIMIT 1)"
					rows := sqlmock.NewRows([]string{"id"}).
						AddRow(false)
					m.ExpectQuery(regexp.QuoteMeta(q)).WillReturnRows(rows)
				}
			},
			strings.Join(internal.Tables, ", "),
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := t.Setup(test.mock)
			got, err := m.Tables()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *MySQLTestSuite) TestMySQL_Exists() {
	var q = "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?"

	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Exists": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(q)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
			nil,
		},
		"Doesn't Exist": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(q)).
					WillReturnError(fmt.Errorf("error"))
			},
			"No database found with the name",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := t.Setup(test.mock)
			err := m.Exists()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(err, test.want)
		})
	}
}

func (t *MySQLTestSuite) TestMySQL_Drop() {
	var q = "DROP DATABASE verbis;"

	tt := map[string]struct {
		mock func(m sqlmock.Sqlmock)
		want interface{}
	}{
		"Exists": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(q)).
					WillReturnResult(sqlmock.NewResult(int64(1), 1))
			},
			nil,
		},
		"Doesn't Exist": {
			func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(q)).
					WillReturnError(fmt.Errorf("error"))
			},
			"Error dropping the database with the name",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			m := t.Setup(test.mock)
			err := m.Drop()
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(err, test.want)
		})
	}
}

const (
	testStdoutValue = "mysqldump"
)

func TestShellProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	_, _ = os.Stdout.WriteString(testStdoutValue)
	os.Exit(0)
}

func (t *MySQLTestSuite) TestMySQL_Dump() {
	cmdSuccess := func(command string, args ...string) *exec.Cmd {
		cs := []string{"-test.run=TestShellProcessSuccess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_TEST_PROCESS=1"}
		return cmd
	}

	tt := map[string]struct {
		path string
		file string
		cmd  func(command string, args ...string) *exec.Cmd
		want interface{}
	}{
		"Success": {
			t.path,
			"test",
			cmdSuccess,
			testStdoutValue,
		},
		"No File": {
			"test",
			"test",
			cmdSuccess,
			"No file or directory with the path",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			execCommand = test.cmd
			path := t.path + string(os.PathSeparator) + "test.sql"

			defer func() {
				execCommand = exec.Command
				_ = os.Remove(path)
			}()

			m := t.Setup(nil)
			err := m.Dump(test.path, test.file)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}

			got, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fail("Error reading path: " + path)
			}

			t.Equal(test.want, string(got))
		})
	}
}

func (t *MySQLTestSuite) TestMySQL_ConnectString() {
	m := t.Setup(nil)
	got := m.connectString()
	want := "root:password@tcp(127.0.0.1:3306)/verbis?tls=false&parseTime=true&multiStatements=true"
	t.Equal(want, got)
}
