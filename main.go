package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vmstelecom/huncoding/src/configuration/logger"
	"github.com/vmstelecom/huncoding/src/controller"
	"github.com/vmstelecom/huncoding/src/controller/routes"
	"github.com/vmstelecom/huncoding/src/model/service"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
