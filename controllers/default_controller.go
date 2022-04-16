package controllers

import (
	"go-gin-mysql-boilerplate/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoutes(c *gin.Context) {
	var res = c.Writer
	helper.Responses(res, http.StatusOK, "Welcome To Rest API Golang CRUD & Authentication", nil)
}
