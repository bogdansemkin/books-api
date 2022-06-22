package model

type Book struct {
	Id int `json:"id" bd:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
}
