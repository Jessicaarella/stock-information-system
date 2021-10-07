package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func GetAll(e echo.Context, db *sqlx.DB) ([]User, error) {
	var users []User

	query := `SELECT id, username, name, email, password, created_at FROM users WHERE deleted_at IS NULL`

	err := db.Select(&users, query)
	if err != nil && err != sql.ErrNoRows {
		return users, err
	}

	if len(users) == 0 {
		return users, sql.ErrNoRows
	}

	return users, nil
}

func Get(e echo.Context, db *sqlx.DB, id int) (User, error) {
	var u User

	query := `SELECT id, username, name, email, password, created_at FROM users WHERE deleted_at IS NULL AND id = ?`

	err := db.Get(&u, query, id)
	if err != nil && err != sql.ErrNoRows {
		return u, err
	}

	if err == sql.ErrNoRows {
		return u, sql.ErrNoRows
	}

	return u, nil
}

func Create(e echo.Context, db *sqlx.DB, user User) error {
	query := `INSERT INTO users (username, name, email, password, created_at) VALUES (?, ?, ?, ?, NOW())`

	r, err := db.Exec(query, user.Username, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = int(id)
	return nil
}

func Update(e echo.Context, db *sqlx.DB, user User, id int) error {
	query := `UPDATE users SET
				username = ?,
				name = ?,
				email = ?,
				password = ?, 
				updated_at = NOW()
			WHERE id = ? AND deleted_at IS NULL`

	r, err := db.Exec(query, user.Username, user.Name, user.Email, user.Password, id)
	if err != nil {
		return err
	}

	num, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func Delete(e echo.Context, db *sqlx.DB, id int) error {
	query := `UPDATE users SET
				deleted_at = NOW()
			WHERE id = ? AND deleted_at IS NULL`

	r, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	num, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return sql.ErrNoRows
	}

	return nil
}
