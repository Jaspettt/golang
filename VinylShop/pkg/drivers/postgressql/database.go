package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PGCFG struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Type     string
}

func PGConnect(cfg *PGCFG) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Username, cfg.Password, cfg.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
