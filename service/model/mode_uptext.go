package models

type UpText struct {
	Text string `json:"text" db:"text" binding:"required"`
}
