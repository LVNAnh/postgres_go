package Routes

import (
	"postgres_go/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Controllers.GetUsers)
	router.POST("/users", Controllers.AddUser)

	return router
}
