package api

import (
	"context"
	"github.com/Cleverse/go-utilities/nullable"
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

func (opts DBOptions) GetPagination(model interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	count, err := opts.DB.NewSelect().Model(model).Count(ctx)
	if err != nil {
		log.Fatal("Inside GetPagination: ", err, "\n")
	}
	totalPage := count / opts.Page
	nextPage := nullable.From(opts.Page)
	if totalPage == opts.Page {
		nextPage.SetNull()
	}
	var pagination = models.Pagination{
		TotalRecords: count,
		TotalPage:    count / opts.Page,
		CurrentPage:  opts.Page,
		NextPage:     nextPage,
	}

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
