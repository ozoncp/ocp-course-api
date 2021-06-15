package utils

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func ConnectToPostgres(ctx context.Context, dataSourceName string) (*sqlx.DB, error) {
	var db *sqlx.DB

	err := backoff.Retry(func() error {
		var err error
		db, err = sqlx.Open("pgx", dataSourceName)
		if err != nil {
			log.Debug().Err(err).Msg("Attempt to open connection to DB failed")
			return err
		}
		err = db.Ping()
		if err != nil {
			log.Debug().Err(err).Msg("Attempt to connect to DB failed")
		}
		return err
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))

	if err != nil {
		return nil, err
	}

	return db, nil
}
