package user

import (
	"database/sql"
	"fmt"

	"API/internal/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type IUserRepo interface {
	Insert(user model.User) error
}

func (u *UserRepo) Insert(user model.User) error {
	stmt := `INSERT INTO users(user_id,email,password,created_at)
	VALUES(?,?,?,datetime('now','localtime'));`
	if _, err := u.db.Exec(stmt, user.UserId, user.Email, user.Password); err != nil {
		return fmt.Errorf("insert into db user - %v", err)
	}
	return nil
}
