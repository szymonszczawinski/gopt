package auth

import (
	"embed"
	"errors"
	"fmt"
	"gosi/coreapi/viewhandlers"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	LoginErrorTemplate     = "public/auth/login_error.html"
	LoginErrorTemplateName = "login_error"
)

var (
	ErrorInvalidSessionToken = errors.New("invalid session token")
	ErrorFailedSaveSession   = errors.New("faield to save session")
)

type authHandler struct {
	viewhandlers.BaseHandler
	authService IAuthService
}

func NewAuthHandler(authService IAuthService, fs embed.FS) *authHandler {
	instance := authHandler{
		BaseHandler: viewhandlers.BaseHandler{
			FileSystem: fs,
		},
		authService: authService,
	}
	return &instance
}
func (handler *authHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	routes.Root().GET("login", handler.login)
	routes.Root().POST("login", handler.loginSubmit)
	routes.Root().GET("logout", handler.logout)

}

func (handler *authHandler) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	viewhandlers.AddCompositeView(r, "login", "public/auth/login.html", viewhandlers.GetSimpleLayouts(), handler.FileSystem)
	return r
}

func (handler authHandler) login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "LOGIN",
	})
}

func (handler authHandler) loginSubmit(c *gin.Context) {
	password := c.PostForm("username")
	username := c.PostForm("password")
	loginResult := handler.authService.login(username, password)
	if !loginResult.Sucess() {
		displayLoginError(loginResult.Error(), c, handler.FileSystem)
		return
	}
	log.Println(loginResult.Data())
	sessionToken := uuid.NewString()
	c.SetCookie("session_token", sessionToken, 120, "", "gosi", false, true)
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
	c.Redirect(http.StatusFound, "/gosi")
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func displayLoginError(err error, c *gin.Context, fs embed.FS) {
	log.Println("Login Error", err.Error())
	tmpl := template.Must(template.ParseFS(fs, LoginErrorTemplate))
	tmplerr := tmpl.ExecuteTemplate(c.Writer, LoginErrorTemplateName, gin.H{"error": fmt.Sprintf("Login ERROR: %v", err.Error())})
	if tmplerr != nil {
		log.Println("TEMPLATE ERROR: ", tmplerr.Error())
	}
}
