package handlers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"gopt/core/domain/auth"
	"gopt/coreapi"
	"log/slog"
	"net/http"

	auth_view "gopt/public/auth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

var (
	ErrorInvalidSessionToken = errors.New("invalid session token")
	ErrorFailedSaveSession   = errors.New("faield to save session")

	users = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
)

type IAuthService interface {
	coreapi.IComponent
	Login(username, password string) coreapi.Result[auth.AuthCredentials]
}

type authHandler struct {
	authService IAuthService
}

func NewAuthHandler(authService IAuthService) authHandler {
	instance := authHandler{
		authService: authService,
	}
	return instance
}

func (handler authHandler) ConfigureRoutes(path string, routes Routes) {
	routes.Root().GET("login", handler.login)
	routes.Root().POST("login", handler.loginSubmit)
	routes.Root().GET("logout", handler.logout)
}

func (handler authHandler) login(c *gin.Context) {
	isHxRequest := c.GetHeader("HX-Request")
	if isHxRequest == "true" {
		auth_view.Login(true).Render(c.Request.Context(), c.Writer)
	} else {
		auth_view.Login(false).Render(c.Request.Context(), c.Writer)
	}
}

func (handler authHandler) loginSubmit(c *gin.Context) {
	slog.Info("login submit")
	password := c.PostForm("username")
	username := c.PostForm("password")
	loginResult := handler.authService.Login(username, password)
	if !loginResult.Sucess() {
		auth_view.LoginError(loginResult.Error().Error()).Render(c.Request.Context(), c.Writer)
		return
	}
	slog.Info("login", "result", loginResult.Data())
	// sessionToken := uuid.NewString()
	sessionToken := sha256.Sum256([]byte(loginResult.Data().Username))
	// c.SetCookie("session_token", sessionToken, 120, "", "gopt", false, true)
	session := sessions.Default(c)
	session.Set(USER_ID, fmt.Sprintf("%x", sessionToken))
	if err := session.Save(); err != nil {
		slog.Error("login error save session", "err", err)
	}
	// c.Redirect(http.StatusFound, "/gopt")

	c.Writer.Header().Add("HX-Redirect", "/gopt")
}

func (handler authHandler) logout(c *gin.Context) {
	// home_view.Home().Render(c.Request.Context(), c.Writer)
	// c.Writer.Header().Add("HX-Redirect", "/gopt")

	// session := sessions.Default(c)
	// user := session.Get(auth.AUTH_KEY)
	// if user == nil {
	// 	// c.HTML(http.StatusBadRequest, "error", gin.H{"error": ErrorInvalidSessionToken})
	// 	errors_view.Error("SUPER ERROR").Render(c.Request.Context(), c.Writer)
	// 	return
	// }
	// session.Delete(auth.AUTH_KEY)
	// if err := session.Save(); err != nil {
	// 	c.HTML(http.StatusInternalServerError, "error", gin.H{"error": ErrorFailedSaveSession})
	// 	return
	// }
	c.Redirect(http.StatusFound, "/gopt")
	// c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func SessionAuth(c *gin.Context) {
	slog.Info("Session AUTH for ", "req", c.Request.URL)
	session := sessions.Default(c)
	if session.Get(USER_ID) == nil {
		c.Redirect(http.StatusFound, "/gopt/login")
		c.Abort()
		return
	}

	c.Next()
}
