package model

import (
	"gorm.io/gorm"
	"time"
)

type BasicModel struct {
	ID        int ` json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" `
}
type Task struct {
	BasicModel
	Text      string     ` json:"text"`
	DDL       *time.Time `json:"ddl" `
	Finished  *bool       `json:"isFinish"`
}

type Tasks []Task