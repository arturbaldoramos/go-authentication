package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"

	database "github.com/arturbaldoramos/go-authentication/pkg/db"
	u "github.com/arturbaldoramos/go-authentication/pkg/utils"
)

type User struct {
	gorm.Model

	ID        string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string `gorm:"not null;type:varchar(150)" json:"name"`
	Email     string `gorm:"unique;not null;type:varchar(100);default:null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if user.Name == "" {
		return u.Message(false, "Name shouldn't be empty"), false
	}
	if user.Email == "" {
		return u.Message(false, "Email shouldn't be empty"), false
	}
	if user.Password == "" {
		return u.Message(false, "Password shouldn't be empty"), false
	}

	return u.Message(true, "success"), true
}

func (user *User) Create() map[string]interface{} {

	//Verify body arguments
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	//Check if email already exist on database
	if err := database.DB.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return u.Message(false, "Email already exists")
	}

	//Encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		u.Message(false, "Internal server error")
		return nil
	}
	user.Password = string(hash)

	//Save user to the database
	database.DB.Create(user)

	//Create a new interface, so we don't return sensitivity information
	userResponse := map[string]interface{}{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"CreatedAt": user.CreatedAt,
		"UpdatedAt": user.UpdatedAt,
	}

	//Formatting response
	resp := u.Message(true, "success")
	resp["user"] = userResponse
	return resp
}

func GetUser(uuid string) *User {
	if uuid == "" {
		return nil
	}

	user := &User{}
	err := database.DB.Where("id = ?", uuid).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}
