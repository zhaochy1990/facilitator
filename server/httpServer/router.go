package httpServer

import "github.com/gin-gonic/gin"

func router() *gin.Engine {
	router := gin.Default()

	router.POST("/task", createTask)

	return router
}
