package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(r *gin.RouterGroup) {
	r.GET("login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "LOGIN",
		})
	})
	r.POST("login", func(c *gin.Context) {
		userCredentials := UserCredentials{
			password: c.PostForm("username"),
			username: c.PostForm("userpassword"),
		}
		validateCredentials(userCredentials)
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "LOGIN",
		})
	})

}

func validateCredentials(userCredentials UserCredentials) {
	panic("unimplemented")
}
