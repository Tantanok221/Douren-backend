package db

import (
	"database/sql"
	"fmt"

	"github.com/tantanok221/douren-backend/internal/helper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	PASSWORD string = helper.GetEnv("PASSWORD")
	USER     string = helper.GetEnv("USER")
	HOST     string = helper.GetEnv("HOST")
	DBNAME   string = helper.GetEnv("DBNAME")
	PORT     string = helper.GetEnv("PORT")
)

func Init() *bun.DB {
	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/postgres", DBNAME, USER, PASSWORD, HOST, PORT)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
