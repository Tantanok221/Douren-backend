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

func GetPagination[T any](opts *DBOptions) models.Pagination {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var model T
	count, err := opts.DB.NewSelect().Model(&model).Count(ctx)
	if err != nil {
		log.Fatal("Inside GetPagination: ", err, "\n")
	}
	totalPage := count / opts.Limit
	nextPage := nullable.From(opts.Page + 1)
	val, _ := nextPage.Get()
	if val > totalPage {
		nextPage.SetNull()
	}
	prevPage := nullable.From(opts.Page - 1)
	val, _ = prevPage.Get()
	if val < 1 {
		prevPage.SetNull()
	}
	var pagination = models.Pagination{
		TotalRecords: count,
		TotalPage:    totalPage,
		CurrentPage:  opts.Page,
		NextPage:     nextPage,
		PreviousPage: prevPage,
	}
	return pagination
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
