package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable         = "users"
	rolesTable         = "roles"
	userRolesTable     = "user_roles"
	filmsTable         = "films"
	hallsTable         = "halls"
	sessionsTable      = "sessions"
	filmSessionstTable = "film_sessions"
	seatsTable         = "seats"
	reservationsTable  = "reservations"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("db is good")

	return db, nil
}
