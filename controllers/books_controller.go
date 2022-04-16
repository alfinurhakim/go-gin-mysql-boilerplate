package controllers

import (
	"fmt"
	"go-gin-mysql-boilerplate/helper"
	"go-gin-mysql-boilerplate/lib/config"
	"go-gin-mysql-boilerplate/repositories"
	"go-gin-mysql-boilerplate/validations"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin" // package used to read the .env file
	"gopkg.in/validator.v2"
)

func GetAllBooks(c *gin.Context) {
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
	repository := repositories.NewRepositoryBooks(db)
	func(booksRepository repositories.BooksRepository) {
		results, err := repository.FindAll(&pagination)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusUnprocessableEntity, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Get data books successfully.", results)
	}(repository)
}

func GetOneBooks(c *gin.Context) {

	// open connection db postgres
	db := config.GetConnection()
	var res = c.Writer

	params := c.Params.ByName("id")
	id, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		panic(err)
	}

	// call repository
	repository := repositories.NewRepositoryBooks(db)
	func(booksRepository repositories.BooksRepository) {
		results, err := repository.FindById(id)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusNotFound, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Get data books detail successfully.", results)
	}(repository)
}

func CreateBooks(c *gin.Context) {

	// open connection db postgres
	db := config.GetConnection()
	var res = c.Writer

	//validation form body for standarize
	var books_validation validations.BooksValidation
	if err := c.ShouldBindJSON(&books_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate(books_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	//created_by
	var params string = c.MustGet("user_id").(string)
	s := strings.Split(params, ".")
	created_by, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}

	// call repository
	repository := repositories.NewRepositoryBooks(db)
	func(booksRepository repositories.BooksRepository) {
		results, err := repository.Save(&books_validation, created_by)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusNotFound, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Create data books detail successfully.", results)
	}(repository)
}

func UpdateBooks(c *gin.Context) {

	// open connection db postgres
	db := config.GetConnection()
	id := c.Params.ByName("id")
	var res = c.Writer

	//validation form body for standarize
	var books_validation validations.BooksValidation
	if err := c.ShouldBindJSON(&books_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate(books_validation); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	//updated_by
	var params string = c.MustGet("user_id").(string)
	s := strings.Split(params, ".")
	updated_by, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}

	//id_books
	id_books, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}

	// call repository
	repository := repositories.NewRepositoryBooks(db)
	func(booksRepository repositories.BooksRepository) {
		results, err := repository.Update(id_books, &books_validation, updated_by)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusNotFound, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Update data books successfully.", results)
	}(repository)

}

func DeleteBooks(c *gin.Context) {

	// open connection db postgres
	db := config.GetConnection()
	var res = c.Writer

	params := c.Params.ByName("id")
	id, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}

	// call repository
	repository := repositories.NewRepositoryBooks(db)
	func(booksRepository repositories.BooksRepository) {
		results, err := repository.Delete(id)
		if err != nil {
			helper.ErrorCustomStatus(res, http.StatusNotFound, err.Error())
			return
		}
		helper.Responses(res, http.StatusOK, "Delete data books successfully.", results)
	}(repository)
}
