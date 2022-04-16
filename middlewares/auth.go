package middlewares

import (
	"fmt"
	"net/http"

	token "go-gin-mysql-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Unathorized",
				"data":    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := token.ExtractTokenID(c)
		fmt.Println("err", err)
		if err != nil {
			fmt.Println("Unathorized")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "Unathorized",
				"data":    nil,
			})
			c.Abort()
			return
		} else {
			fmt.Println("Next Routes")
			c.Set("user_id", user_id) // data for controller
			c.Next()
		}
	}
}
