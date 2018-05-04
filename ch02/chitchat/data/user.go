package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      int
	Email     string
	UserId    string
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1,$2,$3,$4) returning uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.Password)
	return
}
