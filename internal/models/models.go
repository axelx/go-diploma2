package models

import (
	"time"
)

type User struct {
	ID          int
	Login       string
	Password    string
	Jwt         string
	Description string
}

type Entity struct {
	ID                 int
	UserID             int
	Text               string
	BankCard           string
	CreatedAtTimestamp int
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

type JWT string
