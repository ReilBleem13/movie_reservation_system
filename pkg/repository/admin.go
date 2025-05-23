package repository

import (
	"fmt"
	"mrs_project/pkg/models"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{
		db: db,
	}
}

func (r *AdminPostgres) AddFilm(film models.Film) error {
	query := fmt.Sprintf(`INSERT INTO %s (title, description, genre, duration, created_at)
		VALUES ($1, $2, $3, $4, $5)`, filmsTable)

	_, err := r.db.Exec(query, film.Title, film.Description, film.Genre, film.Duration, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *AdminPostgres) DeleteFilm(filmID int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, filmsTable)

	result, err := r.db.Exec(query, filmID)
	if err != nil {
		return fmt.Errorf("failed to delete film: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("film with id: %d doesn't exist", filmID)
	}

	return nil
}

func (r *AdminPostgres) UpdateFilm(filmID int, update models.Film) error {
	setParts := []string{}
	args := map[string]interface{}{"id": filmID}

	if update.Title != nil {
		setParts = append(setParts, "title = :title")
		args["title"] = update.Title
	}

	if update.Description != nil {
		setParts = append(setParts, "description = :description")
		args["description"] = update.Description
	}

	if update.Genre != nil {
		setParts = append(setParts, "genre = :genre")
		args["genre"] = update.Genre
	}

	if update.Duration != nil {
		setParts = append(setParts, "duration = :duration")
		args["duration"] = update.Duration
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", filmsTable, strings.Join(setParts, ", "))
	result, err := r.db.NamedExec(query, args)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("film with id: %d doesn't exist", filmID)
	}

	return nil
}
