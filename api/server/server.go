// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/gin-contrib/location"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

type Server struct {
	*gin.Engine
}

func New(d *deps.Deps) *Server {

	// Force log's color
	gin.ForceConsoleColor()

	// Set mode depending on
	gin.SetMode(gin.ReleaseMode)

	// Remove default gin write
	gin.DefaultWriter = ioutil.Discard

	// New router
	r := gin.Default()

	server := &Server{r}

	// Global middleware
	r.Use(logger.Log())
	r.Use(location.Default())

	r.Use(gin.Recovery())

	// Set up Gzip compression
	server.setupGzip(d)

	// Instantiate the server.
	return server
}

func SendRestart() {
	if proc, err := os.FindProcess(os.Getpid()); err != nil {
		log.Printf("FindProcess: %s", err)
		return
	} else {
		_ = proc.Signal(syscall.Signal(0x1))
	}
}

// ListenAndServe runs Verbis on a given port
// Returns errors.INVALID if the server could not start
func (s *Server) ListenAndServe(port int) error {
	const op = "router.ListenAndServe"

	passedPort := strconv.Itoa(port)

	err := s.Run(":" + passedPort)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: fmt.Sprintf("Could not start Verbis on the port %d", port), Operation: op, Err: err}
	}

	return nil
}

// setupGzip - inits the default gzip compression for the server bt
// looking for the options & setting.
func (s *Server) setupGzip(d *deps.Deps) {
	const op = "router.ListenAndServe"

	options := d.Options

	// Bail if there is no
	if !options.Gzip {
		return
	}

	/// Set the default compression & check options
	compression := gzip.DefaultCompression
	switch options.GzipCompression {
	case "best-compression":
		{
			compression = gzip.BestCompression
			break
		}
	case "best-speed":
		{
			compression = gzip.BestSpeed
			break
		}
	}

	// If the use paths is not set, use the excluded extensions
	// or use the excluded paths.
	if !options.GzipUsePaths {
		if len(options.GzipExcludedExtensions) > 0 {
			s.Use(gzip.Gzip(compression, gzip.WithExcludedExtensions(options.GzipExcludedExtensions)))
			return
		}
	} else {
		if len(options.GzipExcludedPaths) > 0 {
			s.Use(gzip.Gzip(compression, gzip.WithExcludedPaths(options.GzipExcludedPaths)))
			return
		}
	}

	s.Use(gzip.Gzip(compression))
}
