package auth

import (
	"embed"
	"errors"
	"fmt"
	"gosi/core/service"
	iservice "gosi/coreapi/service"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	LoginTemplate   = "public/auth/login.html"
	LoginErrorBlock = "error-login"
)

var authService IAuthService

func AddAuthRoutes(r *gin.RouterGroup, fs embed.FS) {
	authService = mustCreateAuthService()
	r.GET("login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "LOGIN",
		})
	})
	r.POST("login", func(c *gin.Context) {
		credentialsData := CredentialsData{
			password: c.PostForm("username"),
			username: c.PostForm("password"),
		}
		uc, err := authService.login(credentialsData)
		if err != nil {
			displayLoginError(err, c, fs)
			return
		}
		log.Println(uc)
		sessionToken := uuid.NewString()
		c.SetCookie("session_token", sessionToken, 120, "", "gosi", false, true)
		c.HTML(http.StatusOK, "home", gin.H{
			"title": "HOME",
		})
	})

}

func validateCredentials(credentialsData CredentialsData) error {
	return errors.New("ERROR")
}
func displayLoginError(err error, c *gin.Context, fs embed.FS) {
	log.Println("Login Error", err.Error())
	tmpl := template.Must(template.ParseFS(fs, LoginTemplate))
	tmpl.ExecuteTemplate(c.Writer, LoginErrorBlock, gin.H{"error": fmt.Sprintf("Login ERROR: %v", err.Error())})
}

func mustCreateAuthService() IAuthService {

	serviceManager, err := service.GetServiceManager()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	service := serviceManager.MustGetComponent(iservice.ComponentTypeAuthService)
	aService, ok := service.(IAuthService)
	if !ok {
		log.Fatal(err.Error())
		return nil
	}
	return aService
}
