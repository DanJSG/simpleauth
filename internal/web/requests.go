package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type RequestProcessor interface {
	RegisterHandlers()
	RegisterHandler()
}

const messageKey = "message"

var badRequestMessage = map[string]string{messageKey: "Request body cannot be empty."}
var internalServerErrorMessage = map[string]string{messageKey: "An unexpected error occurred."}

func HandleJsonPostError(context *gin.Context, err error) {
	if err == nil {
		return
	} else if errors.Is(err, io.EOF) {
		context.AbortWithStatusJSON(http.StatusBadRequest, badRequestMessage)
	} else {
		context.AbortWithStatusJSON(http.StatusInternalServerError, internalServerErrorMessage)
	}
}
