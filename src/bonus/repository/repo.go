package repository

import (
	"bonus/config"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Repo Querier

func NewSqlRepository(db *sql.DB) Repo {
	return New(db)
}

func NewSqlDatabase(cfg *config.DbConfig) (pool *sql.DB, err error) {
	maxTries := 10

	pool, err = sql.Open("pgx", cfg.Url)
	if err != nil {
		return
	}

	for i := 0; i < maxTries; i++ {
		err = pool.Ping()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		continue
	}

	return
}
