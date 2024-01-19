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
		return u.Message(false, "Name should be on the payload"), false
	}
	if user.Email == "" {
		return u.Message(false, "Email should be on the payload"), false
	}
	if user.Password == "" {
		return u.Message(false, "Password should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		u.Message(false, "Internal server error")
		return nil
	}

	user.Password = string(hash)
	database.DB.Create(user)
	resp := u.Message(true, "success")
	resp["user"] = user
	return resp
}

func GetUser(uuid string) *User {
	user := &User{}
	err := database.DB.First(&user, uuid).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}
