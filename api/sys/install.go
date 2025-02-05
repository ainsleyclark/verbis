// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/verbiscms/verbis/api/database"
	"github.com/verbiscms/verbis/api/database/seeds"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/logger"
	"github.com/verbiscms/verbis/api/store"
	"github.com/verbiscms/verbis/api/store/options"
	"strconv"
)

// Install installs the application. The InstallVerbis struct
// will be validated before installing. The system will
// restart dependant on the arg passed.
func (s *Sys) Install(install domain.InstallVerbis) error {
	// Connect to database
	logger.Info("Attempting to connect to database")
	driver, env, err := s.getDatabase(install.InstallDatabase)
	if err != nil {
		return err
	}
	logger.Info("Successfully connected to the database: " + install.DBDatabase)

	// Migrate
	logger.Info("Migrating database")
	err = driver.Install()
	if err != nil {
		return err
	}

	// Obtain the store
	repository, err := store.New(driver, false)
	if err != nil {
		return err
	}

	// Run the seeds
	logger.Info("Attempting to run seeds")
	seeder := seeds.New(driver, repository)
	err = seeder.Seed()
	if err != nil {
		return err
	}
	logger.Info("Successfully ran seeds")

	// Create the owner
	_, err = repository.User.Create(install.ToUser())
	if err != nil {
		return err
	}

	// Update the options
	err = s.installOptions(repository.Options, install.InstallSite, env.Port())
	if err != nil {
		return err
	}

	// Write to the env
	logger.Info("Attempting to write to env file")
	err = env.Install()
	if err != nil {
		return err
	}
	logger.Info("Successfully wrote to .env file")

	return nil
}

// newDB is an alias for creating a new database.
var newDB = database.New

// getDatabase dials the database and returns a new
// database.Driver, or an error if there was a
// problem connecting.
func (s *Sys) getDatabase(id domain.InstallDatabase) (database.Driver, environment.Env, error) {
	logger.Info("Attempting to connect to database")

	env := environment.Env{
		DbDriver:   database.MySQLDriver,
		DbHost:     id.DBHost,
		DbPort:     id.DBPort,
		DbDatabase: id.DBDatabase,
		DbUser:     id.DBUser,
		DbPassword: id.DBPassword,
	}

	db, err := newDB(&env)
	if err != nil {
		return nil, environment.Env{}, err
	}

	logger.Info("Successfully connected to the database: " + id.DBDatabase)
	return db, env, nil
}

func (s *Sys) installOptions(repo options.Repository, install domain.InstallSite, port int) error {
	err := repo.Update("site_url", "http://localhost:"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	err = repo.Update("site_title", install.SiteTitle)
	if err != nil {
		return err
	}
	err = repo.Update("seo_private", install.Robots)
	if err != nil {
		return err
	}
	return nil
}
