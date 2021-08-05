package request

import "time"

type CreateTask struct {
	Text      string     ` json:"text" binding:"required,max=100"`
	DDL       *time.Time `json:"ddl" `
}

type DeleteTask struct {
	IDs []int `json:"id" binding:"required"`
}
type UpdateTask struct {
	ID       int        `json:"id" binding:"required"`
	Text     string     `json:"text" binding:"omitempty,min=1,max=100"`
	DDL      *time.Time `json:"ddl" `
	Finished *bool       `json:"isFinish" `
}


