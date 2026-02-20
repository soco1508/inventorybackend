package models

type User struct {
	UserID string `db:"user_id" json:"userId"`
	Name   string `db:"name" json:"name"`
	Email  string `db:"email" json:"email"`
}
