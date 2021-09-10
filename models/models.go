package models

type User struct {
	ID int64 `json: "id"`
	Name string `json: "name"`
	Email string `json: "email"`
	password string `json: "password"`
	role string `json: "role"`
}