package api

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/tantanok221/douren-backend/db"
	"github.com/tantanok221/douren-backend/models"
)

func GetAllPrimitiveArtist() []byte {
	var result []map[string]interface{}
	db := db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := db.NewSelect().Model(&models.PrimitiveArtist{}).ColumnExpr("*").Limit(20).Scan(ctx, &result)
	if err != nil {
		log.Fatal("Database Connection Fail", err)
	}
	json, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Json Conversion fail", err)
	}
	return json
}
