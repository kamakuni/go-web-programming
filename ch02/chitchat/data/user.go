package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    string
	CreatedAt time.Time
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1,$2,$3,$4) returning uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", session.Uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}
