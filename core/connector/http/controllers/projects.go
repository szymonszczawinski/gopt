package controllers

import (
	"gosi/core/service"
	"gosi/core/storage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProjectsRoutes(rootRoute *gin.RouterGroup) {
	projectsRoute := rootRoute.Group("/projects")
	projectsRoute.GET("/", getProjects)
}

func getProjects(c *gin.Context) {
	log.Println("getProjetcs")
	serviceManager, err := service.GetServiceManager()
	if err == nil {
		service, err := serviceManager.GetService(storage.ISTORAGESERVICE)
		if err == nil {
			storageService, ok := service.(storage.IStorageService)
			if ok {
				projects := storageService.GetProjects()
				c.JSON(http.StatusOK, gin.H{"data": projects})
			} else {
				log.Println("StorageService has incorrect type")
			}
		} else {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
}
