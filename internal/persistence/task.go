package persistence

import "github.com/zhaochy1990/facilitator/internal/model"

type Persistence interface {
	CreateTask(*model.Task) (*model.Task, error)

	GetTasks() (*model.Task, error)

	GetTaskByName(name string) (*model.Task, error)

	CreateTaskRun(*model.TaskRun) (*model.TaskRun, error)
}
