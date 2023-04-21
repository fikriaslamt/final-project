package app

import (
	"final-project/config"
	"final-project/handler"
	"final-project/repository"
	"final-project/route"
	"final-project/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
