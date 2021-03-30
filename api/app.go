// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||

package api

// SuperAdminString defines if the app (Verbis) is being developed
// or is being packaged out for distribution.
var SuperAdminString = "true"
var SuperAdmin = true

// App defines default values before the the user has defined
// any custom properties by updating the database.
var App = struct {
	Title       string
	Description string
	URL         string
	Logo        string
	Version     string
}{
	Title:       "Verbis",
	Description: "A Verbis website. Publish online, build a business, work from home",
	URL:         "http://127.0.0.1:8080",
	Logo:        "/verbis/images/verbis-logo.svg",
	Version:     "0.0.1",
}

const (
	// The web route of the API.
	APIRoute = "/api/v1"

	DefaultTheme = ""

	// this needs to be dynamic
	Platform = "mac"

	AssetsChannel = 10
	UploadChannel = 10
	ServerChannel = 50
)

var AssetsChan = make(chan int, AssetsChannel)
var UploadChan = make(chan int, UploadChannel)
var ServeChan = make(chan int, ServerChannel)
