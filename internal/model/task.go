package model

type Task struct {
	Name   string   `json:"Name" gorm:"primaryKey"`
	Params []string `json:"Params"`
}
