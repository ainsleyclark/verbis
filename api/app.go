package api

// SuperAdmin defines if the app (Verbis) is being developed
// or is being packaged out for distribution.
var SuperAdminString = "true"
var SuperAdmin = true

// App defines default values before the the user has defined
// any custom properties by updating the database.
var App = struct{
	Title   		string
	Description 	string
	Url 			string
	Logo 			string
	Version         string
}{
	Title: "Verbis",
	Description: "A Verbis website. Publish online, build a business, work from home",
	Url: "http://localhost:8080",
	Logo: "/images/verbis-logo.svg",
	Version: "0.0.1",
}