package render

import (
	"bytes"
	"fmt"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/templates"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

func (r *Render) Page(g *gin.Context) ([]byte, error) {
	const op = "Render.GetPage"

	api.ServeChan <- 1
	defer func() {
		<-api.ServeChan
	}()

	url := g.Request.URL.Path
	post, err := r.store.Posts.GetBySlug(url)
	if err != nil {
		return nil, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No page found with the url: %s", url), Operation: op, Err: err}
	}

	postData, err := r.store.Posts.Format(post)
	if err != nil {
		return nil, &errors.Error{Code: errors.INTERNAL, Message: "Could not format post data", Operation: op, Err: err}
	}

	// Check if the file has been cached
	var foundCache bool
	if r.options.CacheServerAssets {
		var cachedTemplate interface{}
		cachedTemplate, foundCache = cache.Store.Get(url)

		if cachedTemplate != nil && foundCache {
			return cachedTemplate.([]byte), nil
		}
	}

	_, err = g.Cookie("verbis-session")
	if err != nil && post.Status != "published" {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Page not published, or user is not logged in", Operation: op, Err: err}
	}

	pt := "index"
	if post.PageTemplate != "default" {
		pt = r.theme.TemplateDir + "/" + post.PageTemplate
	}

	master := ""
	if post.PageLayout != "default" {
		master = r.theme.LayoutDir + "/" + post.PageLayout
	} else {
		pt = pt + r.theme.FileExtension
	}

	tm := templates.NewManager(g, r.store, &postData, r.config)
	gvFrontend := goview.New(goview.Config{
		Root:         paths.Theme(),
		Extension:    r.theme.FileExtension,
		Master:       master,
		Partials:     []string{},
		Funcs:        tm.GetFunctions(),
		DisableCache: !environment.IsProduction(),
	})

	var b bytes.Buffer
	if err := gvFrontend.RenderWriter(&b, pt, tm.GetData()); err != nil {
		panic(err)
	}

	minified, err := r.minify.MinifyBytes(&b, "text/html")
	if err != nil || minified == nil {
		return b.Bytes(), nil
	}

	go func() {
		if r.options.CacheServerTemplates && !foundCache {
			cache.Store.Set(url, minified, cache.RememberForever)
		}
	}()

	return minified, nil
}
