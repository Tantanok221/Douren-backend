package api

import (
	"context"
	"log"
	"time"

	"github.com/tantanok221/douren-backend/models"
	"github.com/uptrace/bun"
)

type DBOptions struct {
	DB    *bun.DB
	Limit int
	Page  int
}

func (opts DBOptions) GetPagination() {

}

func (opts DBOptions) GetAllPrimitiveArtist() []models.PrimitiveArtist {
	var result []models.PrimitiveArtist
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	query := opts.DB.NewSelect().Model(&result).Limit(opts.Limit)
	if opts.Page > 1 {
		query = query.Offset((opts.Page - 1) * 10)
	}
	err := query.Scan(ctx)
	if err != nil {
		log.Fatal("Database Connection Fail", err)
	}
	return result
}

func GetPrimitiveArtistById(db *bun.DB, id int) []models.PrimitiveArtist {
	var result []models.PrimitiveArtist
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.NewSelect().Model(&result).Where("? = ?", bun.Ident("uuid"), id).Scan(ctx)
	if err != nil {
		log.Fatal("Database Connection Fail", err)
	}
	return result
}
