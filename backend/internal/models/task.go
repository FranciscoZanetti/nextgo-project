package models

type Task struct {
    ID        	string `json:"id"`
    Title     	string `json:"title"`
	Description string `json:"description"`
    CreatedAt 	string `json:"createdAt"`
}
