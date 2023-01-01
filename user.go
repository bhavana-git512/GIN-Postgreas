package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/itsbhavana/gin-gorm-rest/controller"
)

func Userroute( router * gin.Engine){
	router.GET("/",controller.GetController)
	router.POST("/",controller.CreateUser)
	router.DELETE("/",controller.DeleteUser)
	router.PUT("/",controller.PutUser)


	}

