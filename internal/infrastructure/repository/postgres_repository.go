package repository

import (
	"database/sql"
	"errors"
	"tarea-api-usuarios/internal/domain/user"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindAll() ([]user.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *PostgresUserRepository) FindByID(id int) (*user.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	var u user.User
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PostgresUserRepository) Create(u *user.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, u.Name, u.Email).Scan(&u.ID)
}

func (r *PostgresUserRepository) Update(u *user.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	result, err := r.db.Exec(query, u.Name, u.Email, u.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *PostgresUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
