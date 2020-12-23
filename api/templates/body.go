package templates

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// body
//
// Returns class names for the body element. Includes the
// resource, page ID, page title, page template, page
// layout and if the user is logged in or not.
func (t *TemplateFunctions) body() string {
	body := new(bytes.Buffer)
	p := t.post.Post

	// Resource, writes page if no resource (e.g. page)
	if p.Resource == nil {
		body.WriteString("page ")
	} else {
		body.WriteString(fmt.Sprintf("%s ", *p.Resource))
	}

	// Page ID (e.g. page-id-2)
	body.WriteString(fmt.Sprintf("page-id-%v ", p.Id))

	// Title
	body.WriteString(fmt.Sprintf("page-title-%s ", cssValidString(p.Title)))

	// Page template (e.g. page-template-test)
	body.WriteString(fmt.Sprintf("page-template-%s ", cssValidString(p.PageTemplate)))

	// Layout (e.g page-layout-test)
	body.WriteString(fmt.Sprintf("page-layout-%s", cssValidString(p.Layout)))

	// Logged in (e.g. logged-in) if auth
	if t.auth() {
		body.WriteString(" logged-in")
	}

	return body.String()
}

// cssValidString
//
// Strips all special characters from the given string
// and replaces forward slashes with hyphens & spaces
// with dashes for a valid CSS class.
func cssValidString(str string) string {
	r := regexp.MustCompile("[^A-Za-z0-9\\s-/]")

	str = r.ReplaceAllString(str, "")
	str = strings.ReplaceAll(str, "/", "-")
	str = strings.ReplaceAll(str, " ", "-")

	return strings.ToLower(str)
}
