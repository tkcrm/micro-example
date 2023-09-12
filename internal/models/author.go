package models

import "time"

type Author struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CraetedAt time.Time `json:"created_at"`
}
