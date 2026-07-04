package database

import (
	"cleanarchitecture-practice/src/config"
	"context"
	"database/sql"
	"log/slog"
)

type DammyDB struct {
	db *sql.DB
}

func (d *DammyDB) Close() error {
	slog.Info("DammyDBを閉じました")
	return nil
}

var dammyDB *DammyDB

func SetDammyDB(d *DammyDB) {
	dammyDB = d
}

func GetDammyDB() *sql.DB {
	if dammyDB == nil {
		return nil
	}
	return dammyDB.db
}

func NewDammyDB(ctx context.Context, c *config.Config) *DammyDB {
	slog.Info("DammyDBを作成しました")
	d := &DammyDB{db: &sql.DB{}}
	SetDammyDB(d)
	return d
}
