package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/itsbhavana/gin-gorm-rest/config"
	"github.com/itsbhavana/gin-gorm-rest/model"
)
func GetController(c * gin.Context){
	users:=[]modle.User{}
	config.DB.Find(&users)
	c.Json(200,&users)
	
}
func CreateUser(c * gin.Context){
	var user model.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200,&user)

	
}
func DeleteUser(c * gin.Context){
	var user model.User
	config.DB.Where("id=?",c.Param("id")),Delete(&user)
	c.Json(200,&user)
	
}
func PutUser(c * gin.Context){
	var user model.User
	config.DB.Where("id=?",c.Param("id")),First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200,&user)


	
}
