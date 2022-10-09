package persister

import (
	"github.com/zhaochy1990/facilitator/config"
	"github.com/zhaochy1990/facilitator/internal/model"
)

type MongoPersister struct {
	config config.PersisterMongo
}

func NewMongoPersister(c config.PersisterMongo) *MongoPersister {
	return &MongoPersister{
		config: c,
	}
}

func (p MongoPersister) Prepare() error {
	return nil
}

func (MongoPersister) SaveTask(task model.Task) error {
	// TODO implement me
	panic("implement me")
}

func (MongoPersister) GetTaskByID(id string) (*model.Task, error) {
	// TODO implement me
	panic("implement me")
}

func (MongoPersister) SaveTaskDefinition(taskDefinition model.TaskDefinition) error {
	// TODO implement me
	panic("implement me")
}

func (MongoPersister) GetTaskDefinitionByID(id string) (*model.TaskDefinition, error) {
	// TODO implement me
	panic("implement me")
}

func (MongoPersister) SaveAgent(agent model.Agent) error {
	// TODO implement me
	panic("implement me")
}
