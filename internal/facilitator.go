package internal

import (
	"github.com/zhaochy1990/facilitator/config"
	"github.com/zhaochy1990/facilitator/internal/persister"
)

func Start(config config.Config) error {
	p := persister.NewMongoPersister(config.Persister.Mongo)
	err := p.Prepare()
	if err != nil {
		return err
	}

	return nil
}
