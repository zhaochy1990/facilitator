package watcher

import "github.com/go-mysql-org/go-mysql/canal"

type MySQLWatcher struct {
}

func (w *MySQLWatcher) Watch() error {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"
	cfg.Dump.TableDB = "datavault"
	cfg.Dump.Tables = []string{"test"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		return err
	}

	// c.SetEventHandler()

	return c.Run()
}
