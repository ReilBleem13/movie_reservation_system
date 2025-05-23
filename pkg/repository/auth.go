package repository

import (
	"fmt"
	"mrs_project/pkg/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {

	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	err = tx.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (name) VALUE($1)", rolesTable)

	_, err = tx.Exec(query, "user")
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email string) (models.UserWithRole, error) {
	var user models.UserWithRole
	query := fmt.Sprintf(`
		SELECT r.name AS role, u.id, u.name, u.email, u.password_hash, u.created_at
		FROM %s u
		JOIN %s ur ON u.id = ur.user_id
		JOIN %s r ON ur.role_id = r.id
		WHERE u.email = $1`, usersTable, userRolesTable, rolesTable)

	err := r.db.QueryRow(query, email).Scan(&user.Role, &user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return models.UserWithRole{}, err
	}
	return user, nil
}
