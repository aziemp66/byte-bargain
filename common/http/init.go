package http

import (
	"io"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	Router *gin.Engine
}

func NewHTTPServer(ginMode string) HTTPServer {
	if ginMode == "release" {
		gin.SetMode(ginMode)
	}
	if ve, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ve.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
	}

	logFile, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	gin.DisableConsoleColor()
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	return HTTPServer{
		Router: router,
	}
}
