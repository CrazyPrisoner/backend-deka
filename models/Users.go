package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserPSQL struct {
	ID        int64      `json:"id" gorm:PRIMARY_KEY; AUTO_INCREMENT`
	Email     string     `json:"email" gorm:"size:150;not null"`
	Name      string     `json:"name" gorm:"size:150"`
	Password  string     `json:"password,omitempty" gorm:"size:500;not null"`
	StartDate time.Time  `json:"start_date" gorm:"default:CURRENT_TIMESTAMP"`
	EndDate   *time.Time `json:"end_date"`
	HistoryId int64      `json:"history_id"`
}

type JwtClaims struct {
	jwt.StandardClaims
	ID        int64
	Email     string
	Name      string
	StartDate time.Time
	EndDate   *time.Time
	HistoryId int64
}

func CreateUser() string {
	return "Create user"
}

func GetUserByID() string {
	return "Get by ID"
}

func GetAllUsers() string {
	return "Get all users"
}

func DeleteUser() string {
	return "Removing User"
}
