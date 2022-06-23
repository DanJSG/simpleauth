package auth

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"text/template"
)

type API struct {
	Version       int
	APIRootPath   string
	GroupRootPath string
	Routes        []string
	Log           logrus.FieldLogger
}

const apiPathTemplateName = "qualifiedPath"

var apiPathTemplate = newApiPathTemplate()

func newApiPathTemplate() *template.Template {
	tpl, err := template.New(apiPathTemplateName).Parse("{{ .APIRootPath }}/v{{ .Version }}/{{ .GroupRootPath }}")
	if err != nil {
		panic(err)
	}
	return tpl
}

func getApiBasePath(API *API) (string, error) {
	buf := &bytes.Buffer{}
	err := apiPathTemplate.Execute(buf, API)
	if err != nil {
		API.Log.Errorf("Error thrown whilst qualifying base URL. API values: %v. Error: %v", API, err)
		return "", err
	}
	return buf.String(), nil
}

func (API *API) RegisterHandlers(router gin.IRouter) error {
	apiBasePath, err := getApiBasePath(API)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", apiBasePath)
	return nil
}
