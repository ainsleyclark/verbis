package templates

import (
	"encoding/json"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
)

type TemplateFunctions struct {
	gin *gin.Context
	post *domain.Post
	fields map[string]interface{}
	store *models.Store
	functions template.FuncMap
}

// Construct
func NewFunctions(g *gin.Context, s *models.Store, p *domain.Post) *TemplateFunctions {

	f := make(map[string]interface{})
	if p.Fields != nil {
		if err := json.Unmarshal(*p.Fields, &f); err != nil {
			log.Error(err)
		}
	}

	tf := &TemplateFunctions{
		gin: g,
		post: p,
		fields: f,
		store: s,
	}

	return tf
}

// Get all template functions
func (t *TemplateFunctions) GetFunctions() template.FuncMap {

	funcMap := template.FuncMap{
		// Env
		//"appEnv": t.appEnv,
		"isProduction": t.isProduction,
		"isDebug": t.isDebug,
		// Header & Footer
		"verbisHead": t.getHeader,
		"verbisFoot": t.getFooter,
		// Posts
		"getResource": t.getResource,
		// Fields
		"getField": t.getField,
		"getFields": t.getFields,
		"hasField": t.hasField,
		"getRepeater": t.getRepeater,
		"getFlexible": t.getFlexible,
		"getSubField": t.getSubField,
		// Auth
		"isAuth": t.isAuth,
		"isAdmin": t.isAdmin,
		// Posts
		"getPost": t.getPost,
		// Media
		"getMedia": t.getMedia,
		// Paths
		"assets": t.assetsPath,
		"storage": t.storagePath,
		// Partials
		"partial": t.partial,
		// Helpers
		"fullUrl": t.getFullUrl,
		"escape": t.escape,
	}

	t.functions = funcMap

	return funcMap
}

// GetData - Returns all the necessary data for template usage.
func (t *TemplateFunctions) GetData() (map[string]interface{}, error) {

	 theme, err := t.store.Site.GetThemeConfig()
	 if err != nil {
	 	return nil, err
	 }

	 templates, err := t.store.Site.GetTemplates()
	 if err != nil {
		return nil, err
	 }
	 //
	 //var test = make(map[string]map[string]string)
	 //for _, v := range templates.Template {
	 //	 test[v["key"]] = *v
		// fmt.Println(test)
	 //	//layouts[k] = map[string]interface{
	 //	//	v: "e"
		////}
	 //}

	layouts, err := t.store.Site.GetLayouts()
	if err != nil {
		return nil, err
	}
	//var layouts = make(map[string]interface{})
	//for k, v := range l {
	//	l.Layout[k].
	//	layouts[v.] = v
	//}

	 data := map[string]interface{}{
	 	"Site": t.store.Site.GetGlobalConfig(),
	 	"Post": t.post,
	 	"Theme": theme.Theme,
	 	"Resources": theme.Resources,
	 	"Templates": templates,
	 	"Layouts": layouts,
	 }

	return data, nil
}

// Get the app env
func (t *TemplateFunctions) appEnv() string {
	return environment.GetAppEnv()
}

// If the app is in production or development
func (t *TemplateFunctions) isProduction() bool {
	return environment.IsProduction()
}

// If the app is in debug mode
func (t *TemplateFunctions) isDebug() bool {
	return environment.IsDebug()
}

