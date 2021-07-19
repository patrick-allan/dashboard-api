package model

import "time"

type User struct {
	Id         int        `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Email      string     `json:"email,omitempty"`
	Username   string     `json:"username,omitempty"`
	Password   string     `json:"password,omitempty"`
	Birthdate  string     `json:"birthdate,omitempty"`
	Document   string     `json:"document,omitempty"`
	Status     string     `json:"status,omitempty"`
	Created_at *time.Time `json:"created_at,omitempty"`
	Updated_at *time.Time `json:"updated_at,omitempty"`
	Deleted_at *time.Time `json:"deleted_at,omitempty"`
	//Gender      []Gender      `json:"gender,omitempty"`
	//Nationality []Nationality `json:"nationality,omitempty"`
}

type Gender struct {
	Id         int        `json:"id,omitempty"`
	Type       string     `json:"type,omitempty"`
	Status     string     `json:"status,omitempty"`
	Created_at *time.Time `json:"created_at,omitempty"`
	Updated_at *time.Time `json:"updated_at,omitempty"`
	Deleted_at *time.Time `json:"deleted_at,omitempty"`
}

type Nationality struct {
	Id          int        `json:"id,omitempty"`
	Nationality string     `json:"nationality,omitempty"`
	Status      string     `json:"status,omitempty"`
	Created_at  *time.Time `json:"created_at,omitempty"`
	Updated_at  *time.Time `json:"updated_at,omitempty"`
	Deleted_at  *time.Time `json:"deleted_at,omitempty"`
}
