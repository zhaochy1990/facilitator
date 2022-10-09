package scheduler

import "github.com/zhaochy1990/facilitator/internal/model"

type Scheduler struct {
	waitingQueue chan model.Task
}

func NewScheduler(waitingQueue chan model.Task) *Scheduler {
	return &Scheduler{waitingQueue: waitingQueue}
}

func (s *Scheduler) Start() {

}
