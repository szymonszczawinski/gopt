package controllers

import (
	"errors"
	"gosi/core/service"
	"gosi/core/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProjectsRoutes(rootRoute *gin.RouterGroup) {
	projectsRoute := rootRoute.Group("/projects")
	projectsRoute.GET("/", getProjects)
	projectsRoute.GET("/:id", getProject)
}

func getProjects(c *gin.Context) {
	log.Println("getProjetcs")
	storageService, err := getStorageService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": storageService.GetProjects()})
}

func getProject(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect project ID"})
		return
	}

	storageService, err := getStorageService()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	project, err := storageService.GetProject(int64(projectId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project})

}

func getStorageService() (storage.IStorageService, error) {

	serviceManager, err := service.GetServiceManager()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	service, err := serviceManager.GetService(storage.ISTORAGESERVICE)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	storageService, ok := service.(storage.IStorageService)
	if !ok {
		return nil, errors.New("StorageService has incorrect type")
	}
	return storageService, nil
}
