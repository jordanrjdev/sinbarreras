package models

type Subject struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Subjects []Subject
