package models

type Task struct {
	Id       int    `json:"id"`
	Sequence string `json:"sequence"`
	UserId   string `json:"user_id"`
	Priority string `json:"priority"`
}
