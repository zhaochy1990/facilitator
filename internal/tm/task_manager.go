package tm

import (
	"github.com/google/uuid"
	"github.com/zhaochy1990/facilitator/internal/model"
	"github.com/zhaochy1990/facilitator/internal/persister"
)

type TaskManager struct {
	persister persister.Persister

	waitingQueue chan model.Task
}

func NewTaskManager(p persister.Persister, q chan model.Task) *TaskManager {
	return &TaskManager{
		persister:    p,
		waitingQueue: q,
	}
}

// Add enqueued a task to execute
func (t *TaskManager) Add(task model.Task) error {
	// 1. Save this task
	if task.ID == "" {
		task.ID = uuid.New().String()
	}
	err := t.persister.SaveTask(task)
	if err != nil {
		return err
	}
	// 2. Add this task to waiting queue
	t.waitingQueue <- task

	return nil
}
