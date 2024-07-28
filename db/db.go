package db

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/tantanok221/douren-backend/internal/env"
)

var (
	PASSWORD string = env.GetEnv("PASSWORD")
	USER     string = env.GetEnv("USER")
	HOST     string = env.GetEnv("HOST")
	DBNAME   string = env.GetEnv("DBNAME")
	PORT     string = env.GetEnv("PORT")
)

func Init() *bun.DB {
	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/postgres", DBNAME, USER, PASSWORD, HOST, PORT)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}


