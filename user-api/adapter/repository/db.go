package repository

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MaxOpenConns       = 2
	MaxIdleConns       = 2
	ConnMaxLifetimeSec = 100
	ConnMaxIdleSec     = 100
)

func NewRepository() (repo *Repository) {
	//dsn := os.Getenv("POSTGRES_DSN")

	gdb, err := gorm.Open(postgres.Open("user=postuser password=postpass dbname=meli host=db port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to create postgres db: %v", err)
	}

	repo = New(gdb)

	if MaxOpenConns > 0 {
		repo.SetMaxOpenConns(MaxOpenConns)
	}

	if MaxIdleConns > 0 {
		repo.SetMaxIdleConns(MaxIdleConns)
	}

	if ConnMaxLifetimeSec > 0 {
		repo.SetConnMaxLifetime(ConnMaxLifetimeSec)
	}

	if ConnMaxIdleSec > 0 {
		repo.SetConnMaxIdleTime(ConnMaxIdleSec)
	}

	return
}
