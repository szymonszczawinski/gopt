package auth

import (
	"errors"
	"gopt/coreapi/viewhandlers"
	"gopt/public/auth"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrorInvalidSessionToken = errors.New("invalid session token")
	ErrorFailedSaveSession   = errors.New("faield to save session")
)

type authHandler struct {
	authService IAuthService
}

func NewAuthHandler(authService IAuthService) *authHandler {
	instance := authHandler{
		authService: authService,
	}
	return &instance
}

func (handler *authHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	routes.Root().GET("login", handler.login)
	routes.Root().POST("login", handler.loginSubmit)
	routes.Root().GET("logout", handler.logout)
}

func (handler authHandler) login(c *gin.Context) {
	auth.Login().Render(c.Request.Context(), c.Writer)
}

func (handler authHandler) loginSubmit(c *gin.Context) {
	slog.Info("login submit")
	password := c.PostForm("username")
	username := c.PostForm("password")
	loginResult := handler.authService.login(username, password)
	if !loginResult.Sucess() {
		auth.LoginError(loginResult.Error().Error()).Render(c.Request.Context(), c.Writer)
		return
	}
	slog.Info("data", loginResult.Data())
	sessionToken := uuid.NewString()
	c.SetCookie("session_token", sessionToken, 120, "", "gopt", false, true)
	c.HTML(http.StatusOK, "home", gin.H{
		"title": "HOME",
	})
}

func (handler authHandler) logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(AUTH_KEY)
	if user == nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": ErrorInvalidSessionToken})
		return
	}
	session.Delete(AUTH_KEY)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": ErrorFailedSaveSession})
		return
	}
	c.Redirect(http.StatusFound, "/gopt")
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
