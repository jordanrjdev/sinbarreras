package models

type Game struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Subject Subject `json:"subject_id"`
}
