// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"io"
)

// WriterHook is a hook that writes logs of specified
// LogLevels to specified Writer.
type WriterHook struct {
	// The io.Writer, this can be stdout or stderr.
	Writer io.Writer
	// The slice of log levels the writer can too.
	LogLevels []log.Level
}

// Fire
//
// Fire will be called when some logging function is
// called with current hook. It will format log
// entry to string and write it to
// appropriate writer
func (hook *WriterHook) Fire(entry *log.Entry) error {
	const op = "Logger.Hook.Fire"

	line, err := entry.String()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error obtaining the entry string", Operation: op, Err: err}
	}

	_, err = hook.Writer.Write([]byte(line))
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error writing entry to io.Writer", Operation: op, Err: err}
	}

	return nil
}

// Levels
//
// Define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []log.Level {
	return hook.LogLevels
}
