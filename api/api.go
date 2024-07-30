package api

import (
	"context"
	"log"
	"time"

	"github.com/tantanok221/douren-backend/models"
	"github.com/uptrace/bun"
)

func GetAllPrimitiveArtist(db *bun.DB) []models.PrimitiveArtist {
	var result []models.PrimitiveArtist
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.NewSelect().Model(&result).Limit(20).Scan(ctx)
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
