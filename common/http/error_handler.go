package http

import (
	"github.com/gin-gonic/gin"

	errorCommon "github.com/aziemp66/byte-bargain/common/error"
)

func MiddlewareErrorHandler(webURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Set("web_url", webURL)
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]
			// if err can be casted to ClientError, then it is a client error
			if clientError, ok := err.Err.(errorCommon.ClientError); ok {
				c.HTML(clientError.Code, "error", gin.H{
					"error": clientError.Error(),
					"code":  clientError.Code,
				})
			} else if err.IsType(gin.ErrorTypeBind) {
				c.HTML(400, "error", gin.H{
					"error": "Bad request",
					"code":  400,
				})
			} else if err.IsType(gin.ErrorTypePrivate) {
				c.HTML(500, "error", gin.H{
					"error": "Internal server error",
					"code":  500,
				})
			} else {
				c.HTML(500, "error", gin.H{
					"error": "Internal server error",
					"code":  500,
				})
			}
		}
	}
}
