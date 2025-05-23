package models

import "time"

type AboutFilm struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Hall        string    `db:"hall"`
	StartTime   time.Time `db:"start_time"`
	SessionDate time.Time `db:"session_date"`
}

type FreeSeat struct {
	ID      int `db:"id"`
	RowNum  int `db:"row_num"`
	SeatNum int `db:"seat_num"`
}

type UserReservation struct {
	Name        string    `db:"name"`
	Title       string    `db:"title"`
	Hall        string    `db:"hall"`
	StartTime   time.Time `db:"start_time"`
	SessionDate time.Time `db:"session_date"`
	RowNum      int       `db:"row_num"`
	SeatNum     int       `db:"seat_num"`
}

type UserWithRole struct {
	Role      string    `json:"role" db:"role"`
	ID        int       `json:"-" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	Email     string    `json:"email" db:"email" binging:"required"`
	Password  string    `json:"password" db:"password_hash" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
