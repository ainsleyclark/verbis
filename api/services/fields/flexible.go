// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fields

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
	"github.com/verbiscms/verbis/api/logger"
)

// Flexible represents the collection of layouts used
// for the flexible content function in templates.
type Flexible []Layout

// Layout represents the collection of subfield's and
// layout's name.
type Layout struct {
	Name      string
	SubFields SubFields
}

// SubFields represents the collection of fields used
// for templates. It has various functions attached
// to it making it easier to loop over.
type SubFields domain.PostFields

// GetFlexible Returns the collection of Layouts from the
// given key and returns a new Flexible.
// Returns errors.INVALID if the field type is not flexible content.
// Returns errors.INTERNAL if the layouts could not be cast to a string slice.
func (s *Service) GetFlexible(input interface{}, args ...interface{}) Flexible {
	const op = "FieldsService.GetFlexible"

	flexible, ok := input.(Flexible)
	if ok {
		return flexible
	}

	name, err := cast.ToStringE(input)
	if err != nil {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Could not cast input to string", Operation: op, Err: err}).Error()
		return nil
	}

	fields := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return nil
	}

	if field.Type != "flexible" {
		logger.WithError(&errors.Error{Code: errors.INVALID, Message: "Field is not flexible content", Operation: op, Err: fmt.Errorf("field with the name: %s, is not flexible content", name)}).Error()
		return nil
	}

	return s.resolveFlexible("", field, fields)
}

// resolveFlexible Loops over the given layouts (e.g ["layout1","layout2"]
// and builds up an array of SUbFields if the SubField layout
// matches the ranged layout.
// Fields are resolved dependant on the format parameter.
// Returns a new Flexible.
func (s *Service) resolveFlexible(key string, field domain.PostField, fields domain.PostFields) Flexible {
	layouts := field.OriginalValue.Slice()

	var flexible = make(Flexible, len(layouts))
	for index := 0; index < len(flexible); index++ {
		r := walker{
			Key:     key,
			Index:   index,
			Field:   field,
			Fields:  fields,
			Service: s,
		}

		var subFields SubFields
		r.Walk(func(f domain.PostField) {
			subFields = append(subFields, f)
		})

		flexible[index] = Layout{
			Name:      layouts[index],
			SubFields: subFields,
		}
	}

	return flexible
}

// HasRows determines if the Flexible content has any rows.
func (f Flexible) HasRows() bool {
	return len(f) != 0
}

// SubField returns a sub field by key or nil if it wasn't found.
func (s SubFields) SubField(name string) interface{} {
	for _, sub := range s {
		if name == sub.Name {
			return sub.Value
		}
	}
	return nil
}

// First returns the first element in the sub fields, or nil if
// the length of the sub fields is zero.
func (s SubFields) First() interface{} {
	if len(s) == 0 {
		return nil
	}
	return s[0]
}

// Last returns the last element in the sub fields, or nil if
// the length of the sub fields is zero.
func (s SubFields) Last() interface{} {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}
