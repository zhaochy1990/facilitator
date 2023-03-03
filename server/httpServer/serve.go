package httpServer

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zhaochy1990/facilitator/config"
)

func Serve(config config.Config) error {
	port := config.Server.Http.Port
	addr := fmt.Sprintf(":%s", port)
	gin.SetMode(config.Server.Http.Mode)
	router := router()
	return router.Run(addr)
}
