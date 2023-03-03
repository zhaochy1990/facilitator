package server

import (
	"context"

	"github.com/zhaochy1990/facilitator/config"
	"github.com/zhaochy1990/facilitator/server/httpServer"
	"github.com/zhaochy1990/facilitator/server/watcher"
	"github.com/zhaochy1990/x/logger"
	"github.com/zhaochy1990/x/viper"
	"golang.org/x/sync/errgroup"
)

func ServeAll(ctx context.Context, configpath *string) error {
	var configurations config.Config

	viper.MustLoadConfig("FACILITATOR", configpath, &configurations)
	logger.MustGetLogger(&configurations.LoggerConfig)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		logger.S().Infof("starting http server")
		return httpServer.Serve(configurations)
	})

	w := &watcher.MySQLWatcher{}
	g.Go(func() error {
		logger.S().Infof("starting watcher")
		return w.Watch()
	})

	return g.Wait()
}
