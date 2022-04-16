package repositories

import (
	"fmt"
	"go-gin-mysql-boilerplate/models"
	"go-gin-mysql-boilerplate/responses"
	token "go-gin-mysql-boilerplate/utils"
	validations "go-gin-mysql-boilerplate/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CompareUsernamePassword(user_name string, password string, db *gorm.DB) (models.Users, error) {

	var err error

	var u models.Users

	err = db.Where("user_name = ?", user_name).First(&u).Take(&u).Error

	if err != nil {
		return u, err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u, err
	}

	token, err := token.GenerateToken(u.ID)

	fmt.Println("err generate", err)

	if err != nil {
		return u, err
	} else {

		//updated data users to database mysql
		u.Token = string(token)
		db.Save(&u)

		return u, nil
	}

}

//interface
type AuthRepository interface {
	AuthLogin(user_name string, password string) (interface{}, error)
	AuthLogout(bearer_token string) (interface{}, error)
	AuthCurrentUser(c *gin.Context) (interface{}, error)
	AuthRegister(register_validation *validations.RegisterValidation) (interface{}, error)
}

type respositoryAuth struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) *respositoryAuth {
	return &respositoryAuth{db}
}

func (r *respositoryAuth) AuthLogin(user_name string, password string) (interface{}, error) {
	//compare username and password
	token, err := CompareUsernamePassword(user_name, password, r.db)

	if err != nil {
		return nil, err
	}

	//parse data
	var response_login responses.ResponsesLogin
	response_login.DataUser = token
	response_login.Token = token.Token
	return response_login, nil
}

func (r *respositoryAuth) AuthLogout(bearer_token string) (interface{}, error) {
	var users models.Users
	if len(strings.Split(bearer_token, " ")) == 2 {
		bearer_token_new := strings.Split(bearer_token, " ")[1]

		if err := r.db.Where("token = ?", bearer_token_new).First(&users).Error; err != nil {
			return nil, err
		} else {
			//updated data users to database mysql
			users.Token = ""
			r.db.Save(&users)

			return users.Name, nil
		}
	} else {
		return nil, nil
	}
}

func (r *respositoryAuth) AuthCurrentUser(c *gin.Context) (interface{}, error) {
	var users models.Users
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		return nil, err
	}

	//get data by id
	if err := r.db.Where("id = ?", user_id).First(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (r *respositoryAuth) AuthRegister(register_validation *validations.RegisterValidation) (interface{}, error) {

	var users models.Users
	users.Name = register_validation.Name
	users.UserName = register_validation.UserName
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register_validation.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	users.Password = string(hashedPassword)

	//save data to database
	r.db.Create(&users)

	return users, nil
}
