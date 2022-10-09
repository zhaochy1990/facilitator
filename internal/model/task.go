package model

type Task struct {
	ID         string         `json:"ID,omitempty"`
	Name       string         `json:"Name"`
	Params     []string       `json:"Params"`
	TriggerBy  string         `json:"TriggerBy"`
	Properties TaskProperties `json:"Properties"`
}

type TaskProperties struct {
}
