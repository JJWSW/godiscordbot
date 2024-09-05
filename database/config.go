package database

import (
	"alti-radio/common/logger"
	"context"

	"github.com/jackc/pgx/v5"
)

type RadioDatabase struct {
	db *pgx.Conn
	q  *Queries
}

var RadioDB *RadioDatabase

func NewDatabase() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, "user=radio password=radio20240829!! host=127.0.0.1 dbname=radio port=9432 sslmode=disable")
	if err != nil {
		logger.PrintError(5, "Select Error", err)
	}
	q := New(db)
	RadioDB = &RadioDatabase{
		db: db,
		q:  q,
	}
}

func (rd *RadioDatabase) GetQuery() *Queries {
	return rd.q
}
