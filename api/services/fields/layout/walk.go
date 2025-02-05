// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package layout

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/errors"
)

// ByUUID
//
// Traverses the given domain.FieldGroups and compares the
// field UUID until a match has been found.
// Returns a domain.Field if the fields was resolved.
// Returns errors.NOTFOUND if the field was unable to be located or no groups exist.
func ByUUID(uniq uuid.UUID, groups domain.FieldGroups) (domain.Field, error) {
	const op = "Fields.Walker.ByUUID"

	if len(groups) == 0 {
		return domain.Field{}, &errors.Error{Code: errors.NOTFOUND, Message: "No groups exists", Operation: op, Err: fmt.Errorf("no groups exist, unable to range over groups and find fields")}
	}

	for _, g := range groups {
		for _, f := range g.Fields {
			field, found := walkerByUUID(uniq, f)
			if !found {
				continue
			}
			return field, nil
		}
	}

	return domain.Field{}, &errors.Error{Code: errors.NOTFOUND, Message: "Unable to find field", Operation: op, Err: fmt.Errorf("unable to find field with UUID of: %v", uniq)}
}

// ByName
//
// Traverses the given domain.FieldGroups and compares the
// field name until a match has been found.
// Returns a domain.Field if the fields was resolved.
// Returns errors.NOTFOUND if the field was unable to be located or no groups exist.
func ByName(name string, groups domain.FieldGroups) (domain.Field, error) {
	const op = "Fields.Walker.ByUUID"

	if len(groups) == 0 {
		return domain.Field{}, &errors.Error{Code: errors.NOTFOUND, Message: "No groups exists", Operation: op, Err: fmt.Errorf("no groups exist, unable to range over groups and find fields")}
	}

	for _, g := range groups {
		for _, f := range g.Fields {
			field, found := walkerByName(name, f)
			if !found {
				continue
			}
			return field, nil
		}
	}

	return domain.Field{}, &errors.Error{Code: errors.NOTFOUND, Message: "Unable to find field", Operation: op, Err: fmt.Errorf("unable to find field with name of: %s", name)}
}
