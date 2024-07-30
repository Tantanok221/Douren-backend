package api

import (
	"context"
	"log"
	"time"

	"github.com/tantanok221/douren-backend/models"
	"github.com/uptrace/bun"
)


type GetAllPrimitiveArtistOptions struct{
	DB *bun.DB
	Limit int
	Page int
}


func(opts GetAllPrimitiveArtistOptions) GetAllPrimitiveArtist( ) []models.PrimitiveArtist {
	var result []models.PrimitiveArtist
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := opts.DB.NewSelect().Model(&result).Limit(opts.Limit).Scan(ctx)
	if err != nil {
		log.Fatal("Database Connection Fail", err)
	}
	return result
}

func GetPrimitiveArtistById(db *bun.DB, id int) []models.PrimitiveArtist{
	var result []models.PrimitiveArtist
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.NewSelect().Model(&result).Where("? = ?", bun.Ident("uuid"), id).Scan(ctx)
	if err != nil {
		log.Fatal("Database Connection Fail", err)
	}
	return result
}
