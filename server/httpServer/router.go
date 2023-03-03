package httpServer

import "github.com/gin-gonic/gin"

func router() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	v1.POST("/task", createTask)

	return router
}
