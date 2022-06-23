package auth

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestTemplating(t *testing.T) {

	API := &API{
		Log:           logrus.New(),
		APIRootPath:   "api",
		GroupRootPath: "users",
		Version:       1,
	}

	err := API.RegisterHandlers(nil)
	if err != nil {
		return
	}

	println("TEST")

}
