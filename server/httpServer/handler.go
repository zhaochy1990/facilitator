package httpServer

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaochy1990/x/logger"
	"net/http"
)

type CreateTaskRequest struct {
	Type   string `binding:"required" json:"type" form:"type,default=simple"`
	Params []string
}

func createTask(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	logger.S().Infof("create task with type: %s, params: %v", req.Type, req.Params)

	c.JSON(http.StatusOK, gin.H{})
}
