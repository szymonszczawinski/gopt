package auth

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	log.Println("Session AUTH")
	session := sessions.Default(c)
	if session.Get(AUTH_KEY) == nil {
		c.Redirect(http.StatusFound, "/gopt/login")
		c.Abort()
		return
	}

	c.Next()
}
