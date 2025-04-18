package repository

import (
	"database/sql"
	"errors"

	"API/pkg/models"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateAdmin(admin models.Admin) (int, error) {
	var id int
	query := `INSERT INTO admins (email, encrypted_password, created_at, role) VALUES($1, $2, $3, $4) RETURNING id`
	row := r.db.QueryRow(query, admin.Email, admin.Password, admin.CreatedAt, admin.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetAdmin(email string) (*models.Admin, error) {
	var admin models.Admin

	query := `SELECT * FROM admins WHERE email = $1`
	row := r.db.QueryRow(query, email)
	err := row.Scan(&admin.ID, &admin.Email, &admin.Password, &admin.CreatedAt, &admin.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &admin, nil
}

func (r *AuthPostgres) InsertUser(user models.User) (int, error) {
	var id int
	query := `INSERT INTO users (username, email, encrypted_password, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	row := r.db.QueryRow(query, user.UserName, user.Email, user.Password, user.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
