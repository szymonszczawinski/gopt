package handlers

import (
	"errors"
	"gopt/core/domain/auth"
	"gopt/coreapi"
	"log/slog"
	"net/http"

	auth_view "gopt/public/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrorInvalidSessionToken = errors.New("invalid session token")
	ErrorFailedSaveSession   = errors.New("faield to save session")
)

type IAuthService interface {
	coreapi.IComponent
	Login(username, password string) coreapi.Result[auth.AuthCredentials]
}
type authHandler struct {
	authService IAuthService
}

func NewAuthHandler(authService IAuthService) *authHandler {
	instance := authHandler{
		authService: authService,
	}
	return &instance
}

func (handler *authHandler) ConfigureRoutes(routes Routes) {
	routes.Root().GET("login", handler.login)
	routes.Root().POST("login", handler.loginSubmit)
	routes.Root().GET("logout", handler.logout)
}

func (handler authHandler) login(c *gin.Context) {
	auth_view.Login().Render(c.Request.Context(), c.Writer)
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
	slog.Info("data", loginResult.Data())
	sessionToken := uuid.NewString()
	c.SetCookie("session_token", sessionToken, 120, "", "gopt", false, true)
	c.HTML(http.StatusOK, "home", gin.H{
		"title": "HOME",
	})
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
