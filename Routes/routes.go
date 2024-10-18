package Routes

import (
	"postgres_go/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Controllers.GetUsers)
	router.POST("/users", Controllers.AddUser)
	router.PUT("/users/:id", Controllers.UpdateUser)
	router.DELETE("/users/:id", Controllers.DeleteUser)

	router.GET("/cards", Controllers.GetCards)
	router.GET("/cards/:id", Controllers.GetCard)
	router.POST("/cards", Controllers.AddCard)
	router.PUT("/cards/:id", Controllers.UpdateCard)
	router.DELETE("/cards/:id", Controllers.DeleteCard)

	return router
}
