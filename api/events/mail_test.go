// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"bytes"
	"fmt"
	client "github.com/ainsleyclark/go-mail"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/deps"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/environment"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
	site "github.com/verbiscms/verbis/api/mocks/services/site"
	tpl "github.com/verbiscms/verbis/api/mocks/tpl"
	fs "github.com/verbiscms/verbis/api/mocks/verbisfs"
	"github.com/verbiscms/verbis/api/test"
	"github.com/verbiscms/verbis/api/verbisfs"
	"os"
	"testing"
)

// EventTestSuite defines the helper used for event
// testing.
type EventTestSuite struct {
	test.HandlerSuite
	Logger *bytes.Buffer
}

// Assert testing has begun.
func TestEvent(t *testing.T) {
	suite.Run(t, &EventTestSuite{})
}

func (t *EventTestSuite) Setup(withErr bool) *deps.Deps {
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	t.Logger = buf

	mt := &tpl.TemplateHandler{}
	m := &tpl.TemplateExecutor{}

	mt.On("Prepare", mock.Anything).Return(m)

	if withErr {
		m.On("Execute", &bytes.Buffer{}, mock.Anything, mock.Anything).
			Return(mock.Anything, fmt.Errorf("withErr"))
	} else {
		m.On("Execute", &bytes.Buffer{}, mock.Anything, mock.Anything).
			Return(mock.Anything, nil)
	}

	ms := &site.Service{}
	ms.On("Global").Return(domain.Site{})

	d := &deps.Deps{
		Site:    ms,
		Options: &domain.Options{},
	}

	d.SetTmpl(mt)

	return d
}

type mockMailError struct{}

func (m *mockMailError) Send(t *client.Transmission) (client.Response, error) {
	return client.Response{}, fmt.Errorf("error")
}

type mockMailSuccess struct{}

func (m *mockMailSuccess) Send(t *client.Transmission) (client.Response, error) {
	return client.Response{}, nil
}

func (t *EventTestSuite) TestHealthCheck() {
	tt := map[string]struct {
		input *environment.Env
		want  interface{}
	}{
		"Success": {
			&environment.Env{
				MailDriver:      client.SparkPost,
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "name",
				SparkpostAPIKey: "key",
				SparkpostURL:    "https://api.eu.sparkpost.com",
			},
			nil,
		},
		"Error": {
			nil,
			"Environment can't be nil",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := HealthCheck(test.input)
			if got != nil {
				t.Contains(errors.Message(got), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *EventTestSuite) Test_MailValidate() {
	tt := map[string]struct {
		input *mail
		want  interface{}
	}{
		"Success": {
			&mail{
				Client: &mockMailSuccess{},
			},
			nil,
		},
		"Nil Mailer": {
			nil,
			"Mailer is nil",
		},
		"Nil Client": {
			&mail{},
			"Mail client is nil",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got := test.input.Validate()
			if got != nil {
				t.Contains(errors.Message(got), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

func (t *EventTestSuite) TestNew() {
	tt := map[string]struct {
		input *environment.Env
		want  interface{}
	}{
		"Nil Environment": {
			nil,
			"Environment can't be nil",
		},
		"SparkPost": {
			&environment.Env{
				MailDriver:      client.SparkPost,
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "name",
				SparkpostAPIKey: "key",
				SparkpostURL:    "https://api.eu.sparkpost.com",
			},
			client.SparkPost,
		},
		"MailGun": {
			&environment.Env{
				MailDriver:      client.MailGun,
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "name",
				MailGunAPIKey:   "key",
				MailGunURL:      "https://api.eu.sparkpost.com",
				MailGunDomain:   "domain",
			},
			client.MailGun,
		},
		"SendGrid": {
			&environment.Env{
				MailDriver:      client.SendGrid,
				MailFromAddress: "hello@verbiscms.com",
				MailFromName:    "name",
				SendGridAPIKey:  "key",
			},
			client.SendGrid,
		},
		"No Driver": {
			&environment.Env{
				MailDriver: "wrong",
			},
			"No mail driver exists",
		},
		"newMailer Client Error": {
			&environment.Env{
				MailDriver: client.SparkPost,
			},
			"Error validating mail configuration",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			d := &deps.Deps{
				Env: test.input,
			}
			e := event{}
			got, err := newMailer(d, e)
			if err != nil {
				t.Contains(errors.Message(err), test.want)
				return
			}
			t.Equal(test.want, got.Driver)
			t.NotNil(got.Client)
			t.Equal(test.input.MailFromAddress, got.FromAddress)
			t.Equal(test.input.MailFromName, got.FromName)
			t.Equal(e, got.event)
		})
	}
}

func (t *EventTestSuite) Test_MailSend() {
	tt := map[string]struct {
		input mail
		error bool
		mock  func(mock *fs.FS)
		want  interface{}
	}{
		"Success": {
			mail{
				Client: &mockMailSuccess{},
				event: event{
					PlainTextTemplate: "test",
				},
			},
			false,
			func(m *fs.FS) {
				file, err := os.Open(t.T().TempDir())
				t.NoError(err)
				m.On("Open", "mail/test.txt").Return(file, nil)
				m.On("ReadFile", "mail/test.txt").Return([]byte("hello"), nil)
			},
			"Successfully sent email with the subject",
		},
		"Validation Failed": {
			mail{},
			false,
			func(m *fs.FS) {},
			"nil mail client",
		},
		"HTML Error": {
			mail{
				Client: &mockMailSuccess{},
				event:  event{},
			},
			true,
			func(m *fs.FS) {},
			"error",
		},
		"Text Error": {
			mail{
				Client: &mockMailError{},
				event: event{
					PlainTextTemplate: MailDir,
				},
			},
			false,
			func(m *fs.FS) {
				m.On("Open", "mail/mail.txt").Return(nil, fmt.Errorf("error"))
			},
			"error",
		},
		"Send Error": {
			mail{
				Client: &mockMailError{},
				event: event{
					PlainTextTemplate: "test",
				},
			},
			false,
			func(m *fs.FS) {
				file, err := os.Open(t.T().TempDir())
				t.NoError(err)
				m.On("Open", "mail/test.txt").Return(file, nil)
				m.On("ReadFile", "mail/test.txt").Return([]byte("hello"), nil)
			},
			"error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			test.input.Deps = t.Setup(test.error)
			mockFs := &fs.FS{}
			test.mock(mockFs)
			test.input.Deps.FS = &verbisfs.FileSystem{Web: mockFs}
			test.input.Send("data", []string{"hello@verbiscms.com"}, nil)
			t.Contains(t.Logger.String(), test.want)
		})
	}
}
