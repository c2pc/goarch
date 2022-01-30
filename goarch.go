package main

import (
	"github.com/chincharovpc/goarch/cmd"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	cmd.Execute()
}
