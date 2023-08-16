package auth

import (
	"embed"
	"fmt"
	"gosi/coreapi/viewcon"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	LoginTemplate   = "public/auth/login.html"
	LoginErrorBlock = "login-error"
)

type authController struct {
	viewcon.Controller
	authService IAuthService
}

func NewAuthController(authService IAuthService, fs embed.FS) *authController {
	instance := authController{
		Controller: viewcon.Controller{
			FileSystem: fs,
		},
		authService: authService,
	}
	return &instance
}
func (self *authController) ConfigureRoutes(root, pages, api *gin.RouterGroup, fs embed.FS) {
	root.GET("login", self.login)
	root.POST("login", self.loginSubmit)
	root.GET("logout", self.logout)

}

func (self *authController) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	viewcon.AddCompositeTemplate(r, "login", "public/auth/login.html", viewcon.GetSimpleLayouts(), self.FileSystem)
	return r
}

func (self authController) login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "LOGIN",
	})
}

func (self authController) loginSubmit(c *gin.Context) {
	credentialsData := CredentialsData{
		password: c.PostForm("username"),
		username: c.PostForm("password"),
	}
	uc, err := self.authService.login(credentialsData)
	if err != nil {
		displayLoginError(err, c, self.FileSystem)
		return
	}
	log.Println(uc)
	sessionToken := uuid.NewString()
	c.SetCookie("session_token", sessionToken, 120, "", "gosi", false, true)
	c.HTML(http.StatusOK, "home", gin.H{
		"title": "HOME",
	})
}

func (self authController) logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(AUTH_KEY)
	if user == nil {
		c.HTML(http.StatusBadRequest, "error", gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(AUTH_KEY)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "error", gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusFound, "/gosi")
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func displayLoginError(err error, c *gin.Context, fs embed.FS) {
	log.Println("Login Error", err.Error())
	tmpl := template.Must(template.ParseFS(fs, LoginTemplate))
	tmplerr := tmpl.ExecuteTemplate(c.Writer, LoginErrorBlock, gin.H{"error": fmt.Sprintf("Login ERROR: %v", err.Error())})
	if tmplerr != nil {
		log.Println("TEMPLATE ERROR: ", tmplerr.Error())
	}
}
