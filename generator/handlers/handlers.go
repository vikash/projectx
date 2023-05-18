package handlers

import (
	"fmt"
	"github.com/vikash/projectx/generator/config"
	"os"
	"text/template"
)

const temp = `
package handler

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"{{.moduleName}}/store"
)

{{- $Name := .entity.Name | PascalCase }}
{{ $name := .entity.Name | CamelCase -}}

type {{$name}}Handler struct {
    store store.{{$Name}}Store
}

func New{{$Name}}Handler(store store.{{$Name}}Store)  *{{$name}}Handler{
    return &{{$name}}Handler{
        store: store,
    }
}

func (h *{{$name}}Handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.store.Index(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *{{$name}}Handler) Get(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	resp, err := h.store.Retrieve(ctx, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *{{$name}}Handler) Create(ctx *gofr.Context) (interface{}, error) {
	var model store.{{$Name}}

	err := ctx.Bind(&model)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"request body"}}
	}

	resp, err := h.store.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *{{$name}}Handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	err := h.store.Delete(ctx, id)

	return true, err
}

func (h *{{$name}}Handler) Update(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	
	var model store.{{$Name}}
	err := ctx.Bind(&model)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"request body error"}}
	}
	if model.Id != "" {
		return nil, errors.InvalidParam{Param: []string{"request body should not have id."}}
	}
	model.Id = id

	return h.store.Update(ctx, &model)
}
`

func Create(e config.Entity, folderName string, g config.Global, d config.Domain) error {

	file, err := os.Create(folderName + "/" + e.Name + ".go")
	if err != nil {
		return fmt.Errorf("can not create %s. Error: %s", file.Name(), err.Error())
	}
	defer file.Close()

	tmpl := template.Must(template.New("handler").Funcs(config.FuncMap).Parse(temp))
	err = tmpl.ExecuteTemplate(file, "handler", map[string]interface{}{
		"entity":     e,
		"moduleName": g.PackagePrefix + "/" + config.CamelCase(d.Name) + "-service",
	})
	if err != nil {
		return fmt.Errorf("can not parse handler.tmpl. Error: %s", err.Error())
	}

	return nil
}
