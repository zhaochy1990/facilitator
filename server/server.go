package server

import (
	"context"
	"github.com/zhaochy1990/facilitator/config"
	"github.com/zhaochy1990/facilitator/server/httpServer"
	"github.com/zhaochy1990/x/logger"
	"github.com/zhaochy1990/x/viper"
)

func initialize(ctx context.Context, configpath *string) error {
	var configurations config.Config

	viper.MustLoadConfig("FACILITATOR", configpath, &configurations)
	logger.MustGetLogger(&configurations.LoggerConfig)

	return httpServer.Serve(configurations)
}

func ServeAll(ctx context.Context, configpath *string) error {
	return initialize(ctx, configpath)
}
