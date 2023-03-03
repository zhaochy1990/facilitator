package model

type Trigger struct {
	Name   string   `json:"Name" gorm:"primaryKey"`
	Params []string `json:"Params"`
}
