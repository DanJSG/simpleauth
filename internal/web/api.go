package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type API struct {
	Version  int
	RootPath string
	Routes   []string
	Log      *logrus.Logger
}

type Handler interface {
	RegisterHandlers(router gin.IRouter)
}
