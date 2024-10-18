package models

type Task struct {
	Id       int    `json:"id"`
	Sequence string `json:"sequence"`
	UserId   int    `json:"user_id"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
}
