package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	TaskId   string `json:"task_id"`
}
