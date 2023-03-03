package httpServer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhaochy1990/x/logger"
)

type CreateTaskRequest struct {
	Name   string `binding:"required" json:"name"`
	Params []string
}

func createTask(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	logger.S().Infof("create task with name: %s, params: %v", req.Name, req.Params)

	c.JSON(http.StatusOK, gin.H{})
}
