package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"time"
)

const timeout = 10 * time.Second

type ConnectInput struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Connect(input ConnectInput) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     input.Host + ":" + input.Port,
		User:     input.User,
		Password: input.Password,
		Database: input.Database,
		PoolSize: 150,
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
