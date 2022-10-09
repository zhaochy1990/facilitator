package persister

import "github.com/zhaochy1990/facilitator/internal/model"

type Persister interface {
	Prepare() error

	SaveTask(task model.Task) error
	GetTaskByID(id string) (*model.Task, error)

	SaveTaskDefinition(taskDefinition model.TaskDefinition) error
	GetTaskDefinitionByID(id string) (*model.TaskDefinition, error)

	SaveAgent(agent model.Agent) error
}
