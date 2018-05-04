package data

import (
	"time"
)

type Session struct {
	Id        int
	Uuid      int
	Email     string
	UserId    string
	CreatedAt time.Time
}
