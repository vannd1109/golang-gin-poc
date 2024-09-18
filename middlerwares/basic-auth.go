package middlerwares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"vannd": "Vannd.dev3004@",
	})
}
