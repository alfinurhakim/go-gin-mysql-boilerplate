package controllers

import (
	"fmt"
	"go-gin-mysql-boilerplate/helper"
	"go-gin-mysql-boilerplate/lib/config"
	"go-gin-mysql-boilerplate/repositories"
	"net/http"

	"github.com/gin-gonic/gin" // package used to read the .env file
	"gorm.io/gorm"
)

var db *gorm.DB

func GetAllUser(c *gin.Context) {
	// open connection db mysql
	db := config.GetConnection()
	var res = c.Writer

	//handler pagination
	var pagination helper.Pagination
	if c.Bind(&pagination) == nil {
		fmt.Println("====== Bind By Query String ======")
		if pagination.Limit != "" && pagination.Page != "" {
			fmt.Println(pagination.Limit)
			fmt.Println(pagination.Page)
		} else {
			pagination.Limit = "10"
			pagination.Page = "1"
		}
	} else {
		helper.ErrorCustomStatus(res, http.StatusNotFound, "Data Not Found")
		fmt.Println("bind query string nil")
		c.Abort()
		return
	}

	// call repository
	repository := repositories.NewRepositoryUsers(db)
	func(usersRepository repositories.UsersRepository) {
		results, err := repository.FindAll(&pagination)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusUnprocessableEntity, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Get data users successfully.", results)
	}(repository)

}
