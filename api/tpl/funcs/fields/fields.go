package math

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/fields"
	"github.com/ainsleyclark/verbis/api/tpl/internal"
)

// Creates a new fields Namespace
func New(d *deps.Deps, t *internal.TemplateDeps) *Namespace {
	f := fields.NewService(d, *t.Post)
	return &Namespace{
		deps:   d,
		fields: f,
	}
}

// Namespace defines the methods for fields to be used
// as template functions.
type Namespace struct {
	deps   *deps.Deps
	fields fields.FieldService
}

const name = "fields"

// Adds the namespace methods to the internal.FuncsNamespace
// on initialisation.
func Init(d *deps.Deps, t *internal.TemplateDeps) *internal.FuncsNamespace {
	ctx := New(d, t)

	ns := &internal.FuncsNamespace{
		Name: name,
		Context: func(args ...interface{}) interface{} {
			return ctx
		},
	}

	ns.AddMethodMapping(ctx.fields.GetField,
		"field",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetFieldObject,
		"fieldObject",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetFields,
		"fields",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetLayout,
		"layout",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetLayouts,
		"layouts",
		[]string{},
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetRepeater,
		"repeater",
		nil,
		[][2]string{},
	)

	ns.AddMethodMapping(ctx.fields.GetFlexible,
		"flexible",
		nil,
		[][2]string{},
	)

	return ns
}
