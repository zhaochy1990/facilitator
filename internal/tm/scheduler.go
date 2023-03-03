package tm

import "github.com/go-mysql-org/go-mysql/canal"

func Watch() error {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"

}
