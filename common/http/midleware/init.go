package midleware

import (
	sessionCommon "github.com/aziemp66/byte-bargain/common/session"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware(sessionManager *sessionCommon.SessionManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := sessionManager.GetSessionValue(c, "user_id").(string)

		if !ok || userID == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
