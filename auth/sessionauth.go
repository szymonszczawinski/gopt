package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

var (
	users = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
)

func SessionAuth(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(AUTH_KEY) == nil {
		c.Redirect(http.StatusFound, "/gosi/login")
		c.Abort()
		return
	}

	c.Next()
}
