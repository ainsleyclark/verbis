// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolve

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

type ResolverTestSuite struct {
	suite.Suite
}

type noStringer struct{}

func TestResolver(t *testing.T) {
	suite.Run(t, new(ResolverTestSuite))
}

func (t *ResolverTestSuite) BeforeTest(suiteName, testName string) {
	logger.Init()
	logger.SetOutput(ioutil.Discard)
}

func (t *ResolverTestSuite) GetValue() *Value {
	return &Value{
		&deps.Deps{
			Store: &models.Store{},
		},
	}
}
