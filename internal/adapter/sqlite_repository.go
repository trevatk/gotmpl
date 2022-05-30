package adapter

import (
	"context"
	"time"

	"github.com/trevatk/common/database"
	"github.com/trevatk/gotmpl/internal/domain/entity"
)

// SQLiteRepository
type SQLiteRepository struct {
	db *database.DB
}

// ProvideRepository
func ProvideRepository(db *database.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

// Create
func (sr *SQLiteRepository) Create(ctx context.Context, e *entity.Entity) error {

	stmt := ``

	se := sqlEntiy{
		Name: e.Name,
	}

	timeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := sr.db.Conn(timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	tx, err := conn.BeginTx(timeout, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(timeout, stmt, se.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := sr.db.InTx(tx); err != nil {
		return err
	}

	return nil
}
