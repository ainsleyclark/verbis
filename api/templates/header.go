package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"github.com/yosssi/gohtml"
	"html/template"
)


// getHeader obtains all of the site and post wide Code Injection
// as well as any meta information from the page.
func (t *TemplateFunctions) getHeader() template.HTML {
	const op = "Templates.getHeader"

	var b bytes.Buffer

	// Get Code Injection from the Post
	postCi := t.post.CodeInjectHead
	if *postCi != "" {
		b.WriteString(*postCi)
	}

	// Get Code Injection from the Options (globally)
	if t.options.CodeInjectionHead != "" {
		b.WriteString(t.options.CodeInjectionHead)
	}

	// Obtain SEO & set post public
	var seo domain.PostSeo
	postSeo := t.post.SeoMeta.Seo
	postPublic := true
	if postSeo != nil {
		err := json.Unmarshal(*t.post.SeoMeta.Seo, &seo)
		if err != nil {
			log.WithFields(log.Fields{
				"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to unmarshal post seo", Operation: op, Err: err},
			}).Error()
		}
		if !seo.Public {
			postPublic = false
		}
	}

	// Check if the site is public or page is public
	if !t.options.SeoPublic || !postPublic {
		b.WriteString(`<meta name="robots" content="noindex">`)
	}

	// Write the Canonical
	if seo.Canonical != nil && *seo.Canonical != "" {
		b.WriteString(fmt.Sprintf("<link rel=\"canonical\" href=\"%s\" />", *seo.Canonical))
	} else {
		b.WriteString(fmt.Sprintf("<link rel=\"canonical\" href=\"%s\" />", t.site.Url + t.post.Slug))
	}

	// Obtain Meta
	var meta domain.PostMeta
	postMeta := t.post.SeoMeta.Meta
	if postMeta != nil {

		err := json.Unmarshal(*t.post.SeoMeta.Meta, &meta)
		if err != nil {
			log.WithFields(log.Fields{
				"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to unmarshal post meta", Operation: op, Err: err},
			}).Error()
		}

		if meta.Description != "" {
			t.writeMeta(&b, meta.Description)
		} else {
			t.writeMeta(&b, t.options.MetaDescription)
		}

		if meta.Facebook.Title != "" || meta.Facebook.Description != "" {
			t.writeFacebook(&b, meta.Facebook.Title, meta.Facebook.Title, meta.Facebook.ImageId)
		} else {
			t.writeFacebook(&b, t.options.MetaFacebookTitle, t.options.MetaFacebookDescription, t.options.MetaFacebookImageId)
		}

		if meta.Twitter.Title != "" || meta.Twitter.Description != "" {
			t.writeTwitter(&b, meta.Twitter.Title, meta.Twitter.Description, meta.Twitter.ImageId)
		} else {
			t.writeTwitter(&b, t.options.MetaTwitterTitle, t.options.MetaTwitterDescription, t.options.MetaTwitterImageId)
		}

	} else {
		t.writeMeta(&b, t.options.MetaDescription)
		t.writeFacebook(&b, t.options.MetaFacebookTitle, t.options.MetaFacebookDescription, t.options.MetaFacebookImageId)
		t.writeTwitter(&b, t.options.MetaTwitterTitle, t.options.MetaTwitterDescription, t.options.MetaTwitterImageId)
	}


	return template.HTML(gohtml.Format(b.String()))
}


func (t *TemplateFunctions) writeMeta(bytes *bytes.Buffer, description string) {
	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", description))
	}
	bytes.WriteString(fmt.Sprintf("<meta property=\"article:modified_time\" content=\"%s\" />", t.post.PublishedAt))
}


// Facebook
func (t *TemplateFunctions) writeFacebook(bytes *bytes.Buffer, title string, description string, imageId int) {
	if title != "" || description != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:type\" content=\"website\">"))
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:site_name\" content=\"%s\">", t.options.SiteTitle))
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:locale\" content=\"%s\">", t.options.GeneralLocale))
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:type\" content=\"website\" />"))
	}

	if title != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:title\" content=\"%s\">", title))
	}

	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:description\" content=\"%s\">", description))
	}

	image, foundImage := t.store.Media.GetById(imageId)
	if foundImage == nil {
		bytes.WriteString(fmt.Sprintf("<meta property=\"og:image\" content=\"%s\">", t.options.SiteUrl + image.Url))
	}
}

// Twitter
func (t *TemplateFunctions) writeTwitter(bytes *bytes.Buffer, title string, description string, imageId int) {
	if title != "" || description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:card\" content=\"summary\">"))
	}

	if title != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:title\" content=\"%s\">", title))
	}

	if description != "" {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:description\" content=\"%s\">", title))
	}

	image, foundImage := t.store.Media.GetById(imageId)
	if foundImage == nil {
		bytes.WriteString(fmt.Sprintf("<meta name=\"twitter:image\" content=\"%s\">", t.options.SiteUrl + image.Url))
	}
}

// getMetaTitle obtains the meta title from the post, if there is no
// title set on the post, it will look for the global title, if
// none, return empty string.
func (t *TemplateFunctions) getMetaTitle() string {
	const op = "Templates.getMetaTitle"

	var meta domain.PostMeta
	postMeta := t.post.SeoMeta.Meta
	if postMeta != nil {

		err := json.Unmarshal(*t.post.SeoMeta.Meta, &meta)
		if err != nil {
			log.WithFields(log.Fields{
				"error": errors.Error{Code: errors.INTERNAL, Message: "Unable to get options", Operation: op, Err: err},
			}).Error()
		}
	}

	if meta.Title != "" {
		return meta.Title
	}

	if t.options.MetaTitle != "" {
		return t.options.MetaTitle
	}

	return ""
}