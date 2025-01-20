package models

type Chat struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}
