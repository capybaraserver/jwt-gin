package middlewares

import (
	"net/http"

	"jwt-gin/utils/ret"
	"jwt-gin/utils/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, ret.RetFailMsg("Unauthorized"))
			c.Abort()
			return
		}

		claims, err := token.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ret.RetFailMsg("Unauthorized"))
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
