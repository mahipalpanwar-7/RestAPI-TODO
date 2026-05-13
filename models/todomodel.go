package models

type Todo struct{
	Id int  `json:"id"`
	Title string `json:"title"`
	Iscompleted bool `json:"is_completed"`
} 

