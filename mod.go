package suim

import (
	"path"
	"reflect"

	"git.kanosolution.net/kano/kaos"
	"github.com/kanoteknologi/hd"
)

type mod struct{}

func New() *mod {
	return new(mod)
}

func (m *mod) MakeGlobalRoute(svc *kaos.Service) ([]*kaos.ServiceRoute, error) {
	return []*kaos.ServiceRoute{}, nil
}

func (m *mod) MakeModelRoute(svc *kaos.Service, model *kaos.ServiceModel) ([]*kaos.ServiceRoute, error) {
	routes := []*kaos.ServiceRoute{}

	uiModel := model.Model

	//-- form config
	sr := new(kaos.ServiceRoute)
	sr.Path = path.Join(svc.BasePoint(), model.Name, "formconfig")
	sr.Fn = reflect.ValueOf(func(c *kaos.Context, p string) (interface{}, error) {
		cfg, e := CreateFormConfig(uiModel)
		if e == nil && hd.IsHttpHandler(c) {
			hd.SetContentType(c, "application/json")
		}
		return cfg, e
	})
	routes = append(routes, sr)

	//-- grid config
	sr = new(kaos.ServiceRoute)
	sr.Path = path.Join(svc.BasePoint(), model.Name, "gridconfig")
	sr.Fn = reflect.ValueOf(func(c *kaos.Context, p string) (interface{}, error) {
		cfg, e := CreateGridConfig(uiModel)
		if e == nil && hd.IsHttpHandler(c) {
			hd.SetContentType(c, "application/json")
		}
		return cfg, e
	})
	routes = append(routes, sr)

	return routes, nil
}

func (m *mod) Name() string {
	return "suim_ui_config"
}
