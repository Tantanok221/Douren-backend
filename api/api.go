package api

import (
	"context"
	"log"
	"time"

	"github.com/tantanok221/douren-backend/internal/jsonlib"
	"github.com/tantanok221/douren-backend/models"
	"github.com/uptrace/bun"
)

func GetAllPrimitiveArtist(db *bun.DB) []byte {
	var result models.Result
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := models.SelectQuery(db).ColumnExpr("*").Limit(20).Scan(ctx, &result)
	if err != nil {
		log.Fatal("Database Connection Fail", err)
	}
	return jsonlib.GetJson(result)
}
