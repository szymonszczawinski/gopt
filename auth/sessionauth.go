package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// this map stores the users sessions. For larger scale applications, you can use a database or cache for this purpose
var sessions = map[string]session{}

// each session contains the username of the user and the time at which it expires
type session struct {
	username string
	expiry   time.Time
}

// we'll use this method later to determine if the session has expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

type UserCredentials struct {
	password string
	username string
}

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				// c.Writer.WriteHeader(http.StatusUnauthorized)
				c.Redirect(http.StatusFound, "/gosi/login")
				return
			}
			// For any other type of error, return a bad request status
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}

		c.Next()
	}
}
