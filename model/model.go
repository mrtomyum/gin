package model

import "github.com/jmoiron/sqlx"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Secret []byte `json:"secret"`
}

func (u *User) All(db *sqlx.DB) ([]*User, error){
	var users []*User
	sql := `SELECT * FROM user`
	err := db.Select(&users, sql)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) New(db *sqlx.DB) (*User, error) {
	sql := `INSERT INTO user(name) VALUES(?)`
	res, err := db.Exec(sql, u.Name)
	if err != nil {
		return nil, err
	}
	var newUser User
	sql = `SELECT * FROM user WHERE id = ?`
	id, _ := res.LastInsertId()
	err = db.Get(&newUser, sql, id)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

