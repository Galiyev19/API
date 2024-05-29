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
	GetUserByEmail(email string) (model.User, error)
	GetUser(email, password string) (model.User, error)
}

// INSERT USER IN DB
func (u *UserRepo) Insert(user model.User) error {
	stmt := `INSERT INTO users(user_id,email,password,created_at)
	VALUES(?,?,?,datetime('now','localtime'));`
	if _, err := u.db.Exec(stmt, user.UserId, user.Email, user.Password); err != nil {
		return fmt.Errorf("insert into db user - %v", err)
	}
	return nil
}

// GET USER BY EMAIL
func (u *UserRepo) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	stmt := `SELECT * FROM users WHERE email = ?`
	if err := u.db.QueryRow(stmt, email).Scan(&user.UserId, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return model.User{}, fmt.Errorf("NOT FIND USER")
	}
	return user, nil
}

// GET USER
func (u *UserRepo) GetUser(email, password string) (model.User, error) {
	var user model.User
	stmt := `SELECT * FROM users where email = ? AND password = ?`
	if err := u.db.QueryRow(stmt, email).Scan(&user.UserId, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		return model.User{}, fmt.Errorf("NOT FIND USER")
	}
	return user, nil
}
