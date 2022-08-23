package rest

import (
	"goapp"
	"net/http"
)

// AppInterface interface that define the interface of App for Rest adapter
type AppInterface interface {
	goapp.App
	GetRouters() *[]Route
	GetRouterGroup() *[]RouteGroup
	GetMiddlewares() *[]http.Handler
}

// WebserverInterface interface that defines the adapter
type WebserverInterface interface {
	Run()
	Add(app AppInterface)
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
