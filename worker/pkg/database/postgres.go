package database

import (
	"database/sql"

	"github.com/alexandrevilain/image-resizer/worker/pkg/image"
	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	db *sql.DB
}

func Connect(connectionString string) (*PostgresConnection, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &PostgresConnection{
		db: db,
	}, nil
}
func (c *PostgresConnection) Create(image image.ResizeImageJob) error {
	_, err := c.db.Exec(`
		INSERT INTO api.images("desiredWidth", "desiredHeight", "originalUrl", "resultUrl") 
		VALUES($1, $2, $3, $4)
	`, image.Width, image.Height, image.OriginalURL, image.ResultURL)
	return err
}
