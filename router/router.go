// implements routes
package server

import (
	"net/http"

	"go-gin-mysql-boilerplate/controllers"
	"go-gin-mysql-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//handler cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}

func NewRouter(config *viper.Viper) *gin.Engine {
	router := gin.New()

	// Global middleware for blocked cors
	router.Use(CORSMiddleware())

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// base default path
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// define versioning api to /api/v1
	api := router.Group("/api")
	v1 := api.Group("/v1")

	// version 1.0 -------------------------------------------------------

	v1.GET("/", controllers.DefaultRoutes)
	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)

	v1.Use(middlewares.JwtAuthMiddleware())
	{
		v1.GET("/logout", middlewares.CheckAuth(), controllers.Logout)
		v1.GET("/profile", middlewares.CheckAuth(), controllers.CurrentUser)

		//users (Authorized)
		v1.GET("/users", middlewares.CheckAuth(), controllers.GetAllUser)

		//books (Authorized)
		v1.GET("/books", middlewares.CheckAuth(), controllers.GetAllBooks)
		v1.GET("/books/:id", middlewares.CheckAuth(), controllers.GetOneBooks)
		v1.POST("/books", middlewares.CheckAuth(), controllers.CreateBooks)
		v1.PUT("/books/:id", middlewares.CheckAuth(), controllers.UpdateBooks)
		v1.DELETE("/books/:id", middlewares.CheckAuth(), controllers.DeleteBooks)
	}

	// for define documentation swagger
	router.GET("/swagger/api-docs.json", func(c *gin.Context) {
		c.File("docs/swagger.json")
	})
	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/api-docs.json")))

	// version 1.0 -------------------------------------------------------

	return router
}
