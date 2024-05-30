package models

type TodoItem struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
	Done *bool  `json:"done"`
}
