// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/google/uuid"
	"time"
)

type (
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	Post struct {
		Id                int         `db:"id" json:"id" binding:"numeric"`
		UUID              uuid.UUID   `db:"uuid" json:"uuid"`
		Slug              string      `db:"slug" json:"slug" binding:"required,max=150"`
		Title             string      `db:"title" json:"title" binding:"required,max=500"`
		Status            string      `db:"status" json:"status,omitempty"`
		Resource          *string     `db:"resource" json:"resource,max=150"`
		PageTemplate      string      `db:"page_template" json:"page_template,omitempty" binding:"max=150"`
		PageLayout        string      `db:"layout" json:"layout,omitempty" binding:"max=150"`
		CodeInjectionHead *string     `db:"codeinjection_head" json:"codeinjection_head,omitempty"`
		CodeInjectionFoot *string     `db:"codeinjection_foot" json:"codeinjection_foot,omitempty"`
		UserId            int         `db:"user_id" json:"-"`
		PublishedAt       *time.Time  `db:"published_at" json:"published_at"`
		CreatedAt         *time.Time  `db:"created_at" json:"created_at"`
		UpdatedAt         *time.Time  `db:"updated_at" json:"updated_at"`
		SeoMeta           PostOptions `db:"options" json:"options"`
	}
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	//
	Posts []Post
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	//
	PostDatum struct {
		Post     `json:"post"`
		Author   UserPart     `json:"author"`
		Category *Category    `json:"category"`
		Layout   []FieldGroup `json:"layout,omitempty"`
		Fields   PostFields  `json:"fields,omitempty"`
	}
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	//
	PostData []PostDatum
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	//
	PostField struct {
		Id            int         `db:"id" json:"-"`
		PostId        int         `db:"post_id" json:"-"`
		UUID          uuid.UUID   `db:"uuid" json:"uuid" binding:"required"`
		Type          string      `db:"type" json:"type"`
		Name          string      `db:"name" json:"name"`
		Key           string      `db:"field_key" json:"key"`
		Value         interface{} `json:"-"`
		OriginalValue FieldValue  `db:"value" json:"value"`
	}
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	//
	PostFields []PostField
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	//
	PostCreate struct {
		Post
		Author   int         `json:"author,omitempty" binding:"numeric"`
		Category *int        `json:"category,omitempty" binding:"omitempty,numeric"`
		Fields   []PostField `json:"fields,omitempty"`
	}
	// |||||||||||||||||||||||||||||||||||||||||||||||||||||||
	//
	PostOptions struct {
		Id       int       `json:"-"`
		PageId   int       `json:"-" binding:"required|numeric"`
		Meta     *PostMeta `db:"meta" json:"meta"`
		Seo      *PostSeo  `db:"seo" json:"seo"`
		EditLock string    `db:"edit_lock" json:"edit_lock"`
	}
	// PostMeta defines the global meta information for the
	// post used when calling the VerbisHeader.
	PostMeta struct {
		Title       string       `json:"title,omitempty"`
		Description string       `json:"description,omitempty"`
		Twitter     PostTwitter  `json:"twitter,omitempty"`
		Facebook    PostFacebook `json:"facebook,omitempty"`
	}
	// PostTwitter defines the twitter meta information
	// used when calling the VerbisHeader.
	PostTwitter struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		ImageId     int    `json:"image_id,numeric,omitempty"`
	}
	// PostFacebook defines the opengraph meta information
	// used when calling the VerbisHeader.
	PostFacebook struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		ImageId     int    `json:"image_id,numeric,omitempty"`
	}
	// PostSeo defines the options for Seo on the post,
	// including if the post is indexable, if it
	// should appear in the sitemap and any
	// canonical overrides.
	PostSeo struct {
		Public         bool   `json:"public"`
		ExcludeSitemap bool   `json:"exclude_sitemap"`
		Canonical      string `json:"canonical"`
	}
	// PostTemplate defines the Post data for templates when they
	// are used in the front end.
	PostTemplate struct {
		Post
		Author   UserPart
		Category *Category
		Fields   []PostField
	}
)

// Tpl
//
// Converts a PostDatum to a PostTemplate and hides
// layouts.
func (p *PostDatum) Tpl() PostTemplate {
	return PostTemplate{
		Post:     p.Post,
		Author:   p.Author,
		Category: p.Category,
		Fields:   p.Fields,
	}
}

// TypeIsInSlice
//
// Determines if the given field values is in the slice
// passed.
func (f *PostField) TypeIsInSlice(arr []string) bool {
	for _, v := range arr {
		if v == f.Type {
			return true
		}
	}
	return false
}

// Scan
//
// Scanner for PostMeta. unmarshal the PostMeta when
// the entity is pulled from the database.
func (m *PostMeta) Scan(value interface{}) error {
	const op = "Domain.PostMeta.Scan"
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok || bytes == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Scan unsupported for PostMeta", Operation: op, Err: fmt.Errorf("scan not supported")}
	}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling into PostMeta", Operation: op, Err: err}
	}
	return nil
}

// Value
//
// Valuer for PostMeta. marshal the PostMeta when
// the entity is inserted to the database.
func (m *PostMeta) Value() (driver.Value, error) {
	const op = "Domain.PostMeta.Value"
	j, err := json.Marshal(m)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling PostMeta", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}

// Scan
//
// Scanner for PostSeo. unmarshal the PostSeo when
// the entity is pulled from the database.
func (m *PostSeo) Scan(value interface{}) error {
	const op = "Domain.PostSeo.Scan"
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok || bytes == nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Scan unsupported for PostSeo", Operation: op, Err: fmt.Errorf("scan not supported")}
	}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error unmarshalling into PostSeo", Operation: op, Err: err}
	}
	return nil
}

// Value
//
// Valuer for PostSeo. marshal the PostSeo when
// the entity is inserted to the database.
func (m *PostSeo) Value() (driver.Value, error) {
	const op = "Domain.PostSeo.Value"
	j, err := json.Marshal(m)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Error marshalling PostSeo", Operation: op, Err: err}
	}
	return driver.Value(j), nil
}
