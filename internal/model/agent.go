package model

type Agent struct {
	Name    string   `json:"Name" validate:"required"`
	Host    string   `json:"IP" validate:"required"`
	Accepts []string `json:"Accepts" validate:"required,dive,required"` // Refs []TaskDefinition
}
