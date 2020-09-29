package controllers

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/server"
	"github.com/ainsleyclark/verbis/api/templates"
	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// FrontendHandler defines methods for the frontend to interact with the server
type FrontendHandler interface {
	Home(g *gin.Context)
	Test(g *gin.Context)
	GetUploads(g *gin.Context)
	Serve(g *gin.Context)
}

// FrontendController defines the handler for all frontend routes
type FrontendController struct {
	server          *server.Server
	models 			*models.Store
}


// newFrontend - Construct
func newFrontend(m *models.Store) *FrontendController {
	return &FrontendController{
		models: m,
	}
}

// TEMP - Home
func (c *FrontendController) Home(g *gin.Context) {
	g.HTML(200, "templates/home", nil)
}

// TEMP - Test
func (c *FrontendController) Test(g *gin.Context) {
	log.WithFields(log.Fields{
		"error": errors.Error{Code: errors.INTERNAL, Message: "Test", Operation: "op", Err: fmt.Errorf("dfdasfdasf")},
	}).Panic()
	return
}

// GetUploads retrieves images in the uploads folder, returns webp if accepts
func (c *FrontendController) GetUploads(g *gin.Context) {
	acceptHeader := g.Request.Header.Get("Accept")
	acceptWebp := strings.Contains(acceptHeader, "image/webp")

	path := g.Request.URL.Path

	data, mime, err := c.models.Media.Serve(path, acceptWebp)
	if err != nil {
		NoPageFound(g)
		return
	}

	g.Data(200, mime, data)
}

// Serve the front end website
func (c *FrontendController) Serve(g *gin.Context) {

	path := g.Request.URL.Path
	post, err := c.models.Posts.GetBySlug(path)

	if err != nil {
		NoPageFound(g)
		return
	}

	type ResourceData struct {
		Post		domain.Post
	}

	r := ResourceData{
		Post: post,
	}

	tf := templates.NewFunctions(g, c.models, &post)
	gvFrontend := goview.New(goview.Config{
		Root:      paths.Theme(),
		Extension: config.Template.FileExtension,
		Master:    config.Template.LayoutDir + "/" + post.Layout,
		Partials:  []string{},
		Funcs: tf.GetFunctions(),
		DisableCache: !environment.IsProduction(),
	})

	if err := gvFrontend.Render(g.Writer, http.StatusOK, config.Template.TemplateDir + "/" + post.PageTemplate, r); err != nil {
		log.Error(err)
	}
}