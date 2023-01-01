package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsbhavana/gin-gorm-rest/routes"
	"github.com/itsbhavana/gin-gorm-rest/config"
)

func main() {
	router := gin.New()
	config.Connect()
	router.Userroute(router)
	router.Run(":8080")

}
