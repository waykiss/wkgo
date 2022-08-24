package rest

import (
	"github.com/rodrigorodriguescosta/goapp"
	"net/http"
)

// AppInterface interface that define the interface of App for Rest adapter
type AppInterface interface {
	goapp.App
	GetRouters() *[]Route
	GetRouterGroup() *[]RouteGroup
	GetMiddlewares() []func(http.Handler) http.Handler
}

// WebserverInterface interface that defines the adapter
type WebserverInterface interface {
	Run()
	Add(app AppInterface)
	GetApps() []goapp.App
}

type RouteGroup struct {
	Prefix  string
	Routers []Route
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
