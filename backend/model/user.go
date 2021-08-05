package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"time"
)

// BasicModel 由gorm接管,不需要在程序中更改
type BasicModel struct {
	ID        int ` json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" `
}

type User struct {
	BasicModel
	Username string ` json:"username" `
	Password string `json:"-" `
}

type CustomClaims struct {
	ID       int
	Username string
	jwt.StandardClaims
}
