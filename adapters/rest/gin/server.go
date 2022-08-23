package gin

import (
	"github.com/gin-gonic/gin"
	"goapp/adapters/rest"
)

type adapter struct {
	apps []rest.AppInterface
}

var webserver *gin.Engine

var Adapter adapter

//Run method to start fiber webserver
func (f adapter) Run() {
	webserver = gin.Default()
	panic("implement-me")
}
