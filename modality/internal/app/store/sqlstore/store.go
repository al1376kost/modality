package sqlstore

import (
	"database/sql"
	"errors"
	"math"
	"time"

	"modality/internal/app/store"
)

// Store ...
type Store struct {
	db                 *sql.DB
	modalityRepository *ModalityRepository
}

var (
	errNoRows = errors.New("no rows in result set")
)

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Modality ...
func (s *Store) Modality() store.ModalityRepository {
	if s.modalityRepository != nil {
		return s.modalityRepository
	}

	s.modalityRepository = &ModalityRepository{
		store: s,
	}

	return s.modalityRepository
}

// getInt
func (s *Store) getInt(val sql.NullInt64) int {

	if val.Valid {
		return int(val.Int64)
	}
	return 0

}

// getInt64
func (s *Store) getInt64(val sql.NullInt64) int64 {

	if val.Valid {
		return val.Int64
	}
	return 0

}

// getBool
func (s *Store) getBool(val sql.NullBool) bool {

	if val.Valid {
		return val.Bool
	}
	return false

}

// getFloat 2 simbols after point
func (s *Store) getFloat(val sql.NullFloat64, decimal float64) float64 {

	if val.Valid {
		return math.Round(float64(val.Float64)*decimal) / decimal
	}
	return 0.0

}

// getString
func (s *Store) getString(val sql.NullString) string {

	if val.Valid {
		return val.String
	}
	return ""

}

// getTime
func (s *Store) getTime(val sql.NullTime) time.Time {

	if val.Valid {
		return val.Time
	}
	var value time.Time
	return value

}
