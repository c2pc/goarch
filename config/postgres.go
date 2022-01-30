package config

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-pg/pg/v10"
	"time"
)

const timeout = 10 * time.Second

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
}

type PostgresSuperUser struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_SUPERUSER" envDefault:"postgres"`
	Password string `env:"POSTGRES_SUPERUSER_PASSWORD" envDefault:""`
	Database string `env:"POSTGRES_SUPERUSER_DB" envDefault:"postgres"`
}

func GetConnection() (*pg.DB, error) {
	c := GetPostgresConfig()

	db := pg.Connect(&pg.Options{
		Addr:     c.Host + ":" + c.Port,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
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

func GetPostgresConfig() *PostgresConfig {
	c := PostgresConfig{}
	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &c
}

func GetPostgresSuperUserConnection() (*pg.DB, error) {
	c := getPostgresSuperUser()
	db := pg.Connect(&pg.Options{
		Addr:     c.Host + ":" + c.Port,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
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

func getPostgresSuperUser() *PostgresSuperUser {
	c := PostgresSuperUser{}
	if err := env.Parse(&c); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &c
}
