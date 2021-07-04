package model

import "time"

type User struct {
	Id         int       `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Status     string    `json:"status,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}
