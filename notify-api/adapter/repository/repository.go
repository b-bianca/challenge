package repository

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	Conn    *gorm.DB
	connSQL *sql.DB
	Notify  *Notify
}

func New(gdb *gorm.DB) *Repository {
	conn, err := gdb.DB()
	if err != nil {
		return nil
	}

	return &Repository{
		Conn:    gdb,
		connSQL: conn,
		Notify:  NewNotifyRepository(gdb),
	}
}

func (r *Repository) SetMaxOpenConns(n int) {
	r.connSQL.SetMaxOpenConns(n)
}

func (r *Repository) SetMaxIdleConns(n int) {
	r.connSQL.SetMaxIdleConns(n)
}

func (r *Repository) SetConnMaxLifetime(t time.Duration) {
	r.connSQL.SetConnMaxLifetime(t)
}

func (r *Repository) SetConnMaxIdleTime(t time.Duration) {
	r.connSQL.SetConnMaxIdleTime(t)
}
