package repository

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	Conn    *gorm.DB
	ConnSQL *sql.DB
	User    *User
}

func New(gdb *gorm.DB) *Repository {
	conn, err := gdb.DB()
	if err != nil {
		return nil
	}

	return &Repository{
		Conn:    gdb,
		ConnSQL: conn,
		User:    NewUserRepository(gdb),
	}
}

func (r *Repository) SetMaxOpenConns(n int) {
	r.ConnSQL.SetMaxOpenConns(n)
}

func (r *Repository) SetMaxIdleConns(n int) {
	r.ConnSQL.SetMaxIdleConns(n)
}

func (r *Repository) SetConnMaxLifetime(t time.Duration) {
	r.ConnSQL.SetConnMaxLifetime(t)
}

func (r *Repository) SetConnMaxIdleTime(t time.Duration) {
	r.ConnSQL.SetConnMaxIdleTime(t)
}
