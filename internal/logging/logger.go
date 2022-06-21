package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type simpleFormatter struct {
	TimestampFormat string
	LevelNames      []string
}

var defaultFormatter = &simpleFormatter{
	"2006/01/02 - 15:04:05",
	[]string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"},
}

func (formatter *simpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(formatter.TimestampFormat))
	return []byte(fmt.Sprintf("%5s | %s | %s\n", formatter.LevelNames[entry.Level], timestamp, entry.Message)), nil
}

func DefaultLogger() *logrus.Logger {
	return Logger(defaultFormatter)
}

func Logger(formatter logrus.Formatter) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(formatter)
	return log
}

func GinLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}
		logLevel := statusCodeLogLevel(&statusCode)

		log.Logf(logLevel, "%3d | %8v | %15s | %-7s %#v%s",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			errorMessage)
	}
}

func statusCodeLogLevel(code *int) logrus.Level {
	switch {
	case *code >= http.StatusOK && *code < http.StatusBadRequest:
		return logrus.InfoLevel
	case *code >= http.StatusBadRequest && *code < http.StatusInternalServerError:
		return logrus.WarnLevel
	default:
		return logrus.ErrorLevel
	}
}
