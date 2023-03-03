package model

type TaskRun struct {
	ID         string   `json:"ID" gorm:"primaryKey"`
	TaskName   string   `json:"TaskName"`
	Params     []string `json:"Params"`
	StartAt    int32    `json:"StartAt"`
	FinishedAt int32    `json:"FinishedAt"`
	Status     string   `json:"Status"`
}
