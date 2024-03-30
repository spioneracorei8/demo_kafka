package models

import "time"

type Message struct {
	ID          uint      `json:"id"`
	Message     string    `json:"message"`
	Author      string    `json:"author"`
	CreatedDate time.Time `json:"created_date"`
}
