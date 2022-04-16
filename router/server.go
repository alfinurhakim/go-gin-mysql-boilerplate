// start server
package server

import (
	"go-gin-mysql-boilerplate/lib/config"

	"github.com/gin-gonic/gin"
)

// main function for server start
func Init() {
	configEnv := config.GetConfig() // get configuration enviroment ./config/yourenv

	// disable log per production mode
	if configEnv.GetBool("production") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// init router
	r := NewRouter(configEnv)                       // start routes
	r.Run(":" + configEnv.GetString("server.port")) // start server
}
