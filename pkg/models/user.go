package models

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
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

type UserLogin struct {
	Email    string `gorm:"unique;not null;type:varchar(100);default:null" json:"email"`
	Password string `gorm:"not null" json:"password"`
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

	return utils.Message(true, "Success"), true
}

func (user *User) Create() map[string]interface{} {

	//Verify body arguments
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	//Check if email already exist on database, even if user was deleted
	if err := database.DB.Unscoped().Where("email = ?", user.Email).First(&User{}).Error; err == nil {
		return utils.Message(false, "Email already exists")
	}

	//Encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		utils.Message(false, "Internal server error")
		return nil
	}
	user.Password = string(hash)

	//Lower case email
	user.Email = strings.ToLower(user.Email)

	//Save user to the database
	if err := database.DB.Create(user); err == nil {
		return utils.Message(false, "Error saving user")
	}

	//Formatting response
	resp := utils.Message(true, "Success")
	resp["user"] = &userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
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

func GetUserByEmail(email string) *User {
	if email == "" {
		return nil
	}

	user := &User{}
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}

func GetAllUsers() map[string]interface{} {
	var users []*User
	if err := database.DB.Find(&users).Error; err != nil {
		return utils.Message(false, "Error retrieving all users")
	}

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

	resp := utils.Message(true, "Success")
	resp["users"] = simplifiedUsers
	return resp
}

func DeleteUserById(uuid string) map[string]interface{} {
	user := GetUserByID(uuid)

	if user != nil {
		if err := database.DB.Delete(&user).Error; err != nil {
			return utils.Message(false, "Error deleting user")
		}

		resp := utils.Message(true, "Success")
		resp["user"] = &userResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		return resp
	}

	return utils.Message(false, "User not exist")
}

func Login(login *UserLogin) (resp map[string]interface{}, token string) {
	// Get user info
	user := GetUserByEmail(login.Email)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err == nil {

		// Create JWT token
		token := jwt.New(jwt.SigningMethodHS256)

		// Some conversions
		expTimeString := os.Getenv("JWT_EXPIRATION")
		expTime, err := strconv.ParseInt(expTimeString, 10, 64)
		expDuration := time.Duration(expTime)

		// Add info inside token
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Name
		claims["email"] = user.Email
		claims["exp"] = time.Now().Add(time.Hour * expDuration).Unix()

		// Sign the token with the secret key
		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return utils.Message(false, "Failed to generate token"), ""
		}

		// Formatting response
		resp := utils.Message(true, "Success login in")
		return resp, tokenString
	}

	return utils.Message(false, "Username or Password invalid"), ""
}
