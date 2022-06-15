package suim

import (
	"path"
	"reflect"

	"git.kanosolution.net/kano/kaos"
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
	sr.RequestType = reflect.PtrTo(reflect.SliceOf(model.ModelType))
	sr.Fn = reflect.ValueOf(func(*kaos.Context, string) (interface{}, error) {
		cfg, e := CreateFormConfig(uiModel)
		return cfg, e
	})
	routes = append(routes, sr)

	//-- grid config
	sr = new(kaos.ServiceRoute)
	sr.Path = path.Join(svc.BasePoint(), model.Name, "gridconfig")
	sr.RequestType = reflect.PtrTo(reflect.SliceOf(model.ModelType))
	sr.Fn = reflect.ValueOf(func(*kaos.Context, string) (interface{}, error) {
		cfg, e := CreateGridConfig(uiModel)
		return cfg, e
	})
	routes = append(routes, sr)

	return routes, nil
}

func (m *mod) Name() string {
	return "suim_ui_config"
}
