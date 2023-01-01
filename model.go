package model
import(
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"Primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
