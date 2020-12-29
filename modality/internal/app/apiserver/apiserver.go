package apiserver

import (
	"database/sql"

	"modality/internal/app/store/sqlstore"

	_ "github.com/mattn/go-sqlite3"
)

// Start ...
func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)

	srv := newServer(store, config)

	return srv.router.Run(config.BindAddr)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
