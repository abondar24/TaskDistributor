package model

type TypeRequest struct {
	Name   string  `json:"taskName"`
	Status *string `json:"taskStatus"`
}
