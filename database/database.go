package database

import (
	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
)

// Connect to a postgres database
func Connect(url string) *sqlx.DB {
	log := log.WithField("url", url)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.WithError(err).Fatal("failed to open connection to database")
	}
	if err := db.Ping(); err != nil {
		log.WithError(err).Fatal("failed to ping database")
	}
	return db
}
