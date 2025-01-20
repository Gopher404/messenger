package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Chats []Chat `json:"chats"`
}
