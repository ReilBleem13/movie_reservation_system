package repository

import (
	"database/sql"
	"fmt"
	"mrs_project/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ClientagePostgres struct {
	db *sqlx.DB
}

func NewClientagePostgres(db *sqlx.DB) *ClientagePostgres {
	return &ClientagePostgres{
		db: db,
	}
}

func (r *ClientagePostgres) GetFilms() ([]models.AboutFilm, error) {
	var aboutFilms []models.AboutFilm

	query := fmt.Sprintf(`	SELECT f.id, f.title, h.name AS hall, s.start_time, fs.session_date 
							FROM %s fs
							JOIN %s f ON fs.film_id = f.id
							JOIN %s h ON fs.hall_id = h.id
							JOIN %s s ON fs.session_id = s.id
							`, filmSessionstTable, filmsTable, hallsTable, sessionsTable)

	err := r.db.Select(&aboutFilms, query)
	if err != nil {
		return nil, err
	}
	return aboutFilms, nil
}

func (r *ClientagePostgres) GetAvailableSeats(filmSessionID int) ([]models.FreeSeat, error) {
	var freeSeats []models.FreeSeat

	query := fmt.Sprintf(`	SELECT s.id, s.row_num, s.seat_num
							FROM %s s
							LEFT JOIN %s r ON s.id = r.seat_id AND r.film_session_id = $1
							WHERE s.hall_id = (SELECT hall_id FROM film_sessions WHERE id = $1)
							AND r.id IS NULL
							`, seatsTable, reservationsTable)

	err := r.db.Select(&freeSeats, query, filmSessionID)
	if err != nil {
		return []models.FreeSeat{}, err
	}
	return freeSeats, nil
}

func (r *ClientagePostgres) GetReservation(userID int) ([]models.UserReservation, error) {
	var userReservations []models.UserReservation

	query := fmt.Sprintf(`SELECT u.name, f.title, h.name AS hall, s.start_time, fs.session_date, se.row_num, se.seat_num
							FROM %s r
							JOIN %s u ON r.user_id = u.id
							JOIN %s fs ON r.film_session_id = fs.id
							JOIN %s f ON fs.film_id = f.id
							JOIN %s h ON fs.hall_id = h.id
							JOIN %s s ON fs.session_id = s.id
							JOIN %s se ON r.seat_id = se.id
							WHERE u.id = $1`, reservationsTable, usersTable, filmSessionstTable, filmsTable, hallsTable, sessionsTable, seatsTable)

	err := r.db.Select(&userReservations, query, userID)
	if err != nil {
		return []models.UserReservation{}, err
	}
	return userReservations, nil
}

func (r *ClientagePostgres) ReserveSeat(userID, filmSessionID, seatID int) (models.Reservation, error) {
	var reservation models.Reservation

	tx, err := r.db.Beginx()
	if err != nil {
		return models.Reservation{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var filmSessionHallID int
	err = tx.Get(&filmSessionHallID, fmt.Sprintf(`
		SELECT hall_id
		FROM %s
		WHERE id = $1`, filmSessionstTable), filmSessionID)
	if err == sql.ErrNoRows {
		return models.Reservation{}, fmt.Errorf("film session with id %d does not exist", filmSessionID)
	}
	if err != nil {
		return models.Reservation{}, fmt.Errorf("failed to check film session: %w", err)
	}

	var seatHallID int
	err = tx.Get(&seatHallID, fmt.Sprintf(`
		SELECT hall_id 
		FROM %s
		WHERE id = $1`, seatsTable), seatID)
	if err == sql.ErrNoRows {
		return models.Reservation{}, fmt.Errorf("seat with id %d does not exist", seatID)
	}
	if err != nil {
		return models.Reservation{}, fmt.Errorf("failed to check seat: %w", err)
	}

	if filmSessionHallID != seatHallID {
		return models.Reservation{}, fmt.Errorf("seat with %d does not belong to the hall of film session %d", seatID, filmSessionID)
	}

	var existingReservationID int
	err = tx.Get(&existingReservationID, fmt.Sprintf(`
		SELECT id
		FROM %s
		WHERE film_session_id = $1 AND seat_id = $2
		FOR UPDATE`, reservationsTable), filmSessionID, seatID)
	if err == nil {
		return models.Reservation{}, fmt.Errorf("seat with id %d is already reserved for film session %d", seatID, filmSessionID)
	}
	if err != sql.ErrNoRows {
		return models.Reservation{}, fmt.Errorf("failed to check reservation: %w", err)
	}

	err = tx.Get(&reservation,
		fmt.Sprintf(`
			INSERT INTO %s (user_id, film_session_id, seat_id, status)
			VALUES ($1, $2, $3, 'confirmed')
			RETURNING id, user_id, film_session_id, seat_id, status, created_at`, reservationsTable),
		userID, filmSessionID, seatID)
	if err != nil {
		return models.Reservation{}, err
	}

	return reservation, nil
}

func (r *ClientagePostgres) DeleteReservation(userID, filmSessionID, seatID int) error {

	result, err := r.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND film_session_id = $2 AND seat_id = $3", reservationsTable),
		userID, filmSessionID, seatID)
	if err != nil {
		return fmt.Errorf("failed to delete reservation: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no reservation with userID: %d, filmsession:%d, seatID:%d", userID, filmSessionID, seatID)
	}
	return nil
}
