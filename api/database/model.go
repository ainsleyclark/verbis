// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"github.com/verbiscms/verbis/api/common/params"
	strings2 "github.com/verbiscms/verbis/api/common/strings"
	"github.com/verbiscms/verbis/api/database/builder"
	"github.com/verbiscms/verbis/api/errors"
	"regexp"
	"strings"
)

var (
	ErrQueryMessage = "Error executing sql query"
)

// FilterRows takes in the filters from the params set in http.Params
// If there is no filters set, an empty string will be returned.
// Returns errors.INVALID if the operator or column name was not found.
func FilterRows(driver Driver, query *builder.Sqlbuilder, filters map[string][]params.Filter, table string) error {
	const op = "Model.FilterRows"

	operators := []string{"=", ">", ">=", "<", "<=", "<>", "LIKE", "IN", "NOT LIKE", "like", "in", "not like"}

	if len(filters) != 0 {
		counter := 0
		for column, v := range filters {
			// Strip tags
			column = stripAlphaNum(strings.ToLower(column))

			// Check if the column exists before continuing
			var exists bool
			err := driver.DB().QueryRow("SELECT 1 FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? AND COLUMN_NAME = ?", table, column).Scan(&exists)
			if !exists || err != nil {
				return &errors.Error{
					Code:      errors.INVALID,
					Message:   fmt.Sprintf("The %s search query does not exist", column),
					Operation: op,
					Err:       fmt.Errorf("the %s search query does not exists when searching for %s", column, table)}
			}

			var fTable string
			if table != "" {
				fTable = table + "."
			}

			for _, filter := range v {
				// Strip tags
				operator := stripAlphaNum(filter.Operator)
				value := stripAlphaNum(filter.Value)

				// Account for like or not like values
				if operator == "like" || operator == "LIKE" || operator == "not like" || operator == "NOT LIKE" {
					value = "%" + value + "%"
				}

				// Check if the operator exists before continuing
				if opExists := strings2.InSlice(operator, operators); !opExists {
					return &errors.Error{
						Code:      errors.INVALID,
						Message:   fmt.Sprintf("The %s operator does not exist", operator),
						Operation: op,
						Err:       fmt.Errorf("the %s operator does not exists when searching for the %s", operator, fTable)}
				}

				query.WhereRaw(fmt.Sprintf("(%s%s %s '%s')", fTable, column, operator, value))
			}
			counter++
		}
	}

	return nil
}

// stripAlphaNum strips characters and returns an
// alpha numeric string for database processing.
func stripAlphaNum(text string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9 =<>%.@/!+_']+")
	return reg.ReplaceAllString(text, "")
}
