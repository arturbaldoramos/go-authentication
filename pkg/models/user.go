package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"

	database "github.com/arturbaldoramos/go-authentication/pkg/db"
	"github.com/arturbaldoramos/go-authentication/pkg/utils"
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

type userResponse struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string `gorm:"not null;type:varchar(150)" json:"name"`
	Email     string `gorm:"unique;not null;type:varchar(100);default:null" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) Validate() (map[string]interface{}, bool) {
	if user.Name == "" {
		return utils.Message(false, "Name shouldn't be empty"), false
	}
	if user.Email == "" {
		return utils.Message(false, "Email shouldn't be empty"), false
	}
	if user.Password == "" {
		return utils.Message(false, "Password shouldn't be empty"), false
	}

	return utils.Message(true, "success"), true
}

func (user *User) Create() map[string]interface{} {

	//Verify body arguments
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	//Check if email already exist on database
	if err := database.DB.Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return utils.Message(false, "Email already exists")
	}

	//Encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		utils.Message(false, "Internal server error")
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
	resp := utils.Message(true, "success")
	resp["user"] = userResponse
	return resp
}

func GetUserByID(uuid string) *User {
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

func GetAllUsers() map[string]interface{} {
	var users []*User
	database.DB.Find(&users)

	var simplifiedUsers []userResponse
	for _, u := range users {
		simplifiedUsers = append(simplifiedUsers, userResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	resp := utils.Message(true, "success")
	resp["users"] = simplifiedUsers
	return resp
}
