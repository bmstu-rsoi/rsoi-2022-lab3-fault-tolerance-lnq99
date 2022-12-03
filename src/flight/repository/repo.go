package repository

import (
	"database/sql"
	"flight/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

//type IFlightRepo interface {
//	ListFlightsWithOffsetLimit(offset, limit int) []model.Flight
//	//TODO: offset and limit require full table scan, instead use: select * from tb where id>a limit b
//}

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
