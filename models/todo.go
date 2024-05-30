package models

type TodoItem struct {
	Id   int    `json:"id" form:"id"`
	Item string `json:"item" form:"item" binding:"required"`
	Done *bool  `json:"done" form:"done" binding:"required"`
}
