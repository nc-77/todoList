package model

import "time"

type Task struct {
	BasicModel
	Text      string     ` json:"text"`
	DDL       *time.Time `json:"ddl" `
	Finished  *bool       `json:"isFinish"`
	UserId    int        `json:"useId" `
}

type Tasks []Task