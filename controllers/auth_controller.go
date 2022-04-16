package controllers

import (
	"go-gin-mysql-boilerplate/helper"
	"go-gin-mysql-boilerplate/lib/config"
	"go-gin-mysql-boilerplate/repositories"
	"go-gin-mysql-boilerplate/validations"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func Register(c *gin.Context) {
	// open connection db mysql
	db := config.GetConnection()
	var res = c.Writer

	//validation form body for standarize
	var register_validation validations.RegisterValidation
	if err := c.ShouldBindJSON(&register_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := validator.Validate(register_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// call repository
	repository := repositories.NewRepositoryAuth(db)
	func(authRepository repositories.AuthRepository) {
		results, err := repository.AuthRegister(&register_validation)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusUnprocessableEntity, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Register successfully.", results)
	}(repository)
}

func Login(c *gin.Context) {
	// open connection db mysql
	db := config.GetConnection()
	var res = c.Writer

	//validation form body for standarize
	var login_validation validations.LoginValidation
	if err := c.ShouldBindJSON(&login_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	if err := validator.Validate(login_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	// call repository
	repository := repositories.NewRepositoryAuth(db)
	func(authRepository repositories.AuthRepository) {
		results, err := repository.AuthLogin(login_validation.UserName, login_validation.Password)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusForbidden, "username or password is incorrect.")
			return
		}
		helper.Responses(res, http.StatusOK, "Login successfully.", results)
	}(repository)

}

func Logout(c *gin.Context) {
	// open connection db mysql
	db := config.GetConnection()
	var res = c.Writer

	//check for status
	bearerToken := c.Request.Header.Get("Authorization")

	// call repository
	repository := repositories.NewRepositoryAuth(db)
	func(authRepository repositories.AuthRepository) {
		results, err := repository.AuthLogout(bearerToken)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusUnauthorized, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Logout successfully.", results)
	}(repository)
}

func CurrentUser(c *gin.Context) {

	// open connection db mysql
	db := config.GetConnection()
	var res = c.Writer

	// call repository
	repository := repositories.NewRepositoryAuth(db)
	func(authRepository repositories.AuthRepository) {
		results, err := repository.AuthCurrentUser(c)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusUnauthorized, "Unauthorized")
			return
		}
		helper.Responses(res, http.StatusOK, "Get profile successfully.", results)
	}(repository)

}
