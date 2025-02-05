// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import "encoding/json"

// marshaller is the stdlib marshaller as an alias.
var marshaller = json.Marshal

// marshaller is the stdlib marshaller for Valuer
// calls to the DB.
var unmarshaller = json.Unmarshal
