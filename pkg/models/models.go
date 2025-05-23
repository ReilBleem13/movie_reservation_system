package models

import "time"

type User struct {
	ID        int       `json:"-" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	Email     string    `json:"email" db:"email" binging:"required"`
	Password  string    `json:"password" db:"password_hash" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Role struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type UserRole struct {
	UserID int `db:"user_id"`
	RoleID int `db:"role_id"`
}

type Film struct {
	ID          int       `db:"id"`
	Title       *string   `db:"title"`
	Description *string   `db:"description"`
	Genre       *string   `db:"genre"`
	Duration    *int      `db:"duration"`
	CreatedAt   time.Time `db:"created_at"`
}

type Hall struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Capacity  int       `db:"capacity"`
	CreatedAt time.Time `db:"created_at"`
}

type Session struct {
	ID        int       `db:"id"`
	StartTime string    `db:"start_time"` // TIME в PostgreSQL маппится на string
	CreatedAt time.Time `db:"created_at"`
}

type FilmSession struct {
	ID          int       `db:"id"`
	FilmID      int       `db:"film_id"`
	HallID      int       `db:"hall_id"`
	SessionID   int       `db:"session_id"`
	SessionDate time.Time `db:"session_date"` // DATE маппится на time.Time
	TicketPrice float64   `db:"ticket_price"`
	CreatedAt   time.Time `db:"created_at"`
}

type Seat struct {
	ID        int       `db:"id"`
	HallID    int       `db:"hall_id"`
	RowNum    int       `db:"row_num"`
	SeatNum   int       `db:"seat_num"`
	CreatedAt time.Time `db:"created_at"`
}

type Reservation struct {
	ID            int       `db:"id"`
	UserID        int       `db:"user_id"`
	FilmSessionID int       `db:"film_session_id"`
	SeatID        int       `db:"seat_id"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
}
