// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/forms"
	"github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/ainsleyclark/verbis/api/mail/events"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// FormRepository defines methods for Posts to interact with the database
type FormRepository interface {
	Get(meta params.Params) ([]domain.Form, int, error)
	GetById(id int) (domain.Form, error)
	GetByUUID(uuid string) (domain.Form, error)
	//GetValidation(form *domain.Form) dynamicstruct.Builder
	Create(f *domain.Form) (domain.Form, error)
	Update(f *domain.Form) (domain.Form, error)
	Delete(id int) error
	Send(form *domain.Form, ip string, agent string) error
}

// FormsStore defines the data layer for Forms
type FormsStore struct {
	db     *sqlx.DB
	config config.Configuration
	siteModel SiteRepository
}

// newSeoMeta - Construct
func newForms(db *sqlx.DB, config config.Configuration) *FormsStore {
	return &FormsStore{
		db:     db,
		config: config,
		siteModel: newSite(db, config),
	}
}

// Get all forms
//
// Returns errors.INTERNAL if the SQL query was invalid.
// Returns errors.NOTFOUND if there are no forms available.
func (s *FormsStore) Get(meta params.Params) ([]domain.Form, int, error) {
	const op = "FormsRepository.Get"

	var f []domain.Form
	q := fmt.Sprintf("SELECT * FROM forms")
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM forms")

	// Apply filters to total and original query
	filter, err := filterRows(s.db, meta.Filters, "forms")
	if err != nil {
		return nil, -1, err
	}
	q += filter
	countQ += filter

	// Apply order
	q += fmt.Sprintf(" ORDER BY forms.%s %s", meta.OrderBy, meta.OrderDirection)

	// Apply pagination
	if !meta.LimitAll {
		q += fmt.Sprintf(" LIMIT %v OFFSET %v", meta.Limit, (meta.Page-1)*meta.Limit)
	}

	// Select forms
	if err := s.db.Select(&f, q); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get forms", Operation: op, Err: err}
	}

	// Return not found error if no forms are available
	if len(f) == 0 {
		return []domain.Form{}, -1, &errors.Error{Code: errors.NOTFOUND, Message: "No forms available", Operation: op}
	}

	// Count the total number of forms
	var total int
	if err := s.db.QueryRow(countQ).Scan(&total); err != nil {
		return nil, -1, &errors.Error{Code: errors.INTERNAL, Message: "Could not get the total number of forms", Operation: op, Err: err}
	}

	//for _, v := range f {
	//forms.ToStruct(&v)
	//}

	return f, total, nil
}

// GetById - Get the form by Id
//
// Returns errors.NOTFOUND if the form was not found by the given ID.
func (s *FormsStore) GetById(id int) (domain.Form, error) {
	const op = "FormsRepository.GetByUUID"

	var f domain.Form
	if err := s.db.Get(&f, "SELECT * FROM forms WHERE id = ? LIMIT 1", id); err != nil {
		return domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the form with the ID: %v", id), Operation: op, Err: err}
	}

	fields, err := s.GetFields(f.Id)
	if err == nil {
		f.Fields = fields
	}

	//f.Body = forms.ToStruct(&f)

	return f, nil
}

// GetByUUID returns a form by UUID.
//
// Returns errors.NOTFOUND if the form was not found by the given UUID.
func (s *FormsStore) GetByUUID(uuid string) (domain.Form, error) {
	const op = "FormsRepository.GetByUUID"

	var f domain.Form
	if err := s.db.Get(&f, "SELECT * FROM forms WHERE uuid = ? LIMIT 1", uuid); err != nil {
		return domain.Form{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the form with the UUID: %s", uuid), Operation: op, Err: err}
	}

	fields, err := s.GetFields(f.Id)
	if err == nil {
		f.Fields = fields
	}

	f.Body = forms.ToStruct(f)

	return f, nil
}

// GetFields returns form fields by form ID.
//
// Returns errors.NOTFOUND if there were no fields found by the given form ID.
func (s *FormsStore) GetFields(id int) ([]domain.FormField, error) {
	const op = "FormsRepository.GetFields"
	var f []domain.FormField
	if err := s.db.Select(&f, "SELECT * FROM form_fields WHERE form_id = ?", id); err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("Could not get the form fields with the form ID: %v", id), Operation: op, Err: err}
	}
	return f, nil
}

// Create a new form
//
// Returns errors.CONFLICT if the the form (name) already exists.
// Returns errors.INTERNAL if the SQL query was invalid or the function could not get the newly created ID.
func (s *FormsStore) Create(f *domain.Form) (domain.Form, error) {
	const op = "FormsRepository.Create"

	e, err := s.db.Exec("INSERT INTO forms (uuid, name, email_send, email_message, email_subject, store_db, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())", uuid.New().String(), f.Name, f.EmailSend, f.EmailMessage, f.EmailSubject, f.StoreDB)
	if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the form with the name: %v", f.Name), Operation: op, Err: err}
	}

	id, err := e.LastInsertId()
	if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not get the newly created form ID with the name: %v", f.Name), Operation: op, Err: err}
	}
	f.Id = int(id)

	for _, v := range f.Fields {
		_, err := s.db.Exec("INSERT INTO form_fields (uuid, form_id, key, label, type, validation, required, options, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())", uuid.New().String(), f.Id, v.Key, v.Type, v.Validation, v.Required, v.Options)
		if err != nil {
			return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the form fields with the key: %v", v.Key), Operation: op, Err: err}
		}
	}

	nf, err := s.GetById(int(id))
	if err != nil {
		return domain.Form{}, err
	}

	return nf, nil
}

// Update category
//
// Returns errors.NOTFOUND if the form was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FormsStore) Update(f *domain.Form) (domain.Form, error) {
	const op = "FormsRepository.Update"

	_, err := s.GetById(f.Id)
	if err != nil {
		return domain.Form{}, err
	}

	_, err = s.db.Exec("UPDATE forms SET name = ?, email_send = ?, email_message = ?, email_subject = ?, store_db = ?, updated_at = NOW() WHERE id = ?", f.Name, f.EmailSend, f.EmailMessage, f.EmailSubject, f.StoreDB, f.Id)
	if err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not update the form with the name: %s", f.Name), Operation: op, Err: err}
	}

	if _, err := s.db.Exec("DELETE FROM form_fields WHERE form_id = ?", f.Id); err != nil {
		return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete form fields with the form ID: %v", f.Id), Operation: op, Err: err}
	}

	for _, v := range f.Fields {
		_, err := s.db.Exec("INSERT INTO form_fields (uuid, form_id, key, label, type, validation, required, options, updated_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())", uuid.New().String(), f.Id, v.Key, v.Type, v.Validation, v.Required, v.Options)
		if err != nil {
			return domain.Form{}, &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the form fields with the key: %v", v.Key), Operation: op, Err: err}
		}
	}

	return *f, nil
}

// Delete form from forms and form fields table
//
// Returns errors.NOTFOUND if the category was not found.
// Returns errors.INTERNAL if the SQL query was invalid.
func (s *FormsStore) Delete(id int) error {
	const op = "FormsRepository.Delete"

	_, err := s.GetById(id)
	if err != nil {
		return err
	}

	if _, err := s.db.Exec("DELETE FROM forms WHERE id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete the form with the ID: %v", id), Operation: op, Err: err}
	}

	if _, err := s.db.Exec("DELETE FROM form_fields WHERE form_id = ?", id); err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not delete form fields with the form ID: %v", id), Operation: op, Err: err}
	}

	return nil
}

func (s *FormsStore) Send(form *domain.Form, ip string, agent string) error {
	const op = "FormsRepository.GetFields"

	fv, att, err := forms.NewReader(form).Values()
	if err != nil {
		return err
	}

	if form.StoreDB {
		err := s.storeSubmission(form, fv, ip, agent)
		if err != nil {
			return err
		}
	}

	if form.EmailSend {
		err := s.mailSubmission(form, fv, att)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *FormsStore) mailSubmission(form *domain.Form, values forms.FormValues, attachments forms.Attachments) error {
	const op = "FormsRepository.mailSubmission"
	fs, err := events.NewFormSend(s.config)
	if err != nil {
		return err
	}
	if err := fs.Send(&events.FormSendData{
		Site:   s.siteModel.GetGlobalConfig(),
		Form:   form,
		Values: values,
	}, attachments); err != nil {
		return err
	}
	return nil
}

func (s *FormsStore) storeSubmission(form *domain.Form, values forms.FormValues, ip string, agent string) error {
	const op = "FormsRepository.storeSubmission"

	f, err := values.JSON()
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not process the form fields for storing"), Operation: op, Err: err}
	}

	_, err = s.db.Exec("INSERT INTO form_submissions (uuid, form_id, fields, ip_address, user_agent, sent_at) VALUES (?, ?, ?, ?, ?, NOW())", uuid.New().String(), form.Id, f, ip, agent)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: fmt.Sprintf("Could not create the form submission with the ID: %v", form.Id), Operation: op, Err: err}
	}

	return nil
}
