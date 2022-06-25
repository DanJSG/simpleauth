package auth

import (
	"fmt"
	"github.com/danjsg/simpleauth/internal/collections"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v *Version) FullyQualified() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) MajorMinor() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

type API struct {
	Version       Version
	BasePath      string
	RelativePaths collections.Set[string]
	Log           logrus.FieldLogger
}

func (API *API) RegisterHandlers(router gin.IRouter) {
	API.RegisterHandlersWithMiddleware(router)
}

func (API *API) RegisterHandlersWithMiddleware(router gin.IRouter, middleware ...gin.HandlerFunc) {
	basePath := fmt.Sprintf("/v%s/%s", API.Version.MajorMinor(), API.BasePath)
	baseRouter := router.Group(basePath, middleware...)
	API.registerAllHandlers(baseRouter)
}

func noOpHandler(context *gin.Context) {}

func (API *API) registerAllHandlers(router gin.IRouter) {
	// TODO add handlers
	API.post(router, "/authorize", noOpHandler)
	API.post(router, "/", noOpHandler)
	API.get(router, "/token", noOpHandler)
	API.post(router, "/token/revoke", noOpHandler)
}

func (API *API) get(router gin.IRouter, route string, handlerFunc gin.HandlerFunc) {
	API.register(router, route, "GET", handlerFunc)
}

func (API *API) post(router gin.IRouter, route string, handlerFunc gin.HandlerFunc) {
	API.register(router, route, "POST", handlerFunc)
}

func (API *API) put(router gin.IRouter, route string, handlerFunc gin.HandlerFunc) {
	API.register(router, route, "PUT", handlerFunc)
}

func (API *API) delete(router gin.IRouter, route string, handlerFunc gin.HandlerFunc) {
	API.register(router, route, "DELETE", handlerFunc)
}

func (API *API) register(router gin.IRouter, route string, method string, handlerFunc gin.HandlerFunc) {
	router.Handle(method, route, handlerFunc)
	API.RelativePaths.Add(route)
}
