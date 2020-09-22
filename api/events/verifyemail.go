package events

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/encryption"
	"github.com/ainsleyclark/verbis/api/mail"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type VerifyEmail struct {
	mailer *mail.Mailer
}

// Create a new verify email event.
func NewVerifyEmail() (*VerifyEmail, error) {

	m, err := mail.New()
	if err != nil {
		return &VerifyEmail{}, err
	}

	return &VerifyEmail{
		mailer: m,
	}, nil
}

// Send the verify email event.
func (e *VerifyEmail) Send(u *domain.User, title string) error {

	md5String := encryption.MD5Hash(strconv.Itoa(u.Id) + u.Email)

	data := mail.Data{
		"AppUrl": environment.GetAppName(),
		"AppTitle": title,
		"AdminPath": config.Admin.Path,
		"Token": md5String,
		"UserName": u.FirstName,
	}

	html, err := e.mailer.ExecuteHTML("verify-email.html", data)
	if err != nil {
		log.Error(err)
	}

	tm := mail.Sender{
		To:      	[]string{u.Email},
		Subject: 	"Thanks for signing up " + u.FirstName,
		HTML: 		html,
	}

	_, err = e.mailer.Send(&tm)
	if err != nil {
		return err
	}

	return nil
}
