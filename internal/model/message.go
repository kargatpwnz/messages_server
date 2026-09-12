package model

import "time"

type Message struct {
	ID        uint64    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at,omitzero"`
}
