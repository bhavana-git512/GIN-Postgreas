package config
import (
	"github.com/itsbhavana/gin-gorm-rest/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var db *gorm.DB
func Connect(){
	
	db, err:= gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),&gorm.Config{})

	if err != nil {
		panic(err)

	}
	db.AutoMigrate(&model.User{})
	DB=db

}
