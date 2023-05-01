package midleware

import (
	"github.com/gin-gonic/gin"

	sessionCommon "github.com/aziemp66/byte-bargain/common/session"
)

func SessionAuthMiddleware(sessionManager *sessionCommon.SessionManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := sessionManager.GetSessionValue(c, "user_id").(string)

		if !ok || userID == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"code":  401,
				"error": "Unauthorized",
			})
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
