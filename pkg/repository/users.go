package repository

import (
	"API/pkg/models"
	"database/sql"
	"errors"
)

type UsersPostgres struct {
	db *sql.DB
}

func NewUsersPostgres(db *sql.DB) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

func (r *UsersPostgres) GetUserList() (*[]models.User, error) {
	query := `SELECT * FROM users ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (r *UsersPostgres) GetUserByID(ID int) (*models.User, error) {
	var user models.User

	query := `SELECT id, username, email, encrypted_password, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, ID).Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersPostgres) UpdateUser(user models.User, ID int) error {
	query := `UPDATE users SET username = $1, email = $2, encrypted_password = $3 WHERE id = $4`

	_, err := r.db.Exec(query, user.UserName, user.Email, user.Password, ID)
	if err != nil {
		return errors.New("failed to update user: " + err.Error())
	}

	return nil
}
