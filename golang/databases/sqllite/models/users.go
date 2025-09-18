package models

import (
	"context"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

type UserModel struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

func NewUsersDb(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) Insert(user UserModel) (*UserModel, error) {
	res, err := u.db.ExecContext(context.Background(), "INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	nUser := &UserModel{
		Id:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return nUser, nil

}

func (u *Users) Get(id int64) (*UserModel, error) {
	var res UserModel
	err := u.db.QueryRowContext(context.Background(), "SELECT * FROM users WHERE id = ?", id).Scan(&res.Id, &res.Name, &res.Email, &res.Password)
	if err != nil {
		return nil, err
	}
	fmt.Print(res)
	return &res, nil
}

func (u *Users) List() (users []*UserModel, err error) {
	res, err := u.db.QueryContext(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer res.Close()
	for res.Next() {
		t := &UserModel{}
		res.Scan(&t.Id, &t.Name, &t.Email, &t.Password)
		users = append(users, t)
	}
	return 
}
