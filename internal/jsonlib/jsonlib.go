package jsonlib

import (
	"encoding/json"
	"log"

	"github.com/tantanok221/douren-backend/models"
)

func GetJson(result models.Result) []byte {
	json, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Json Conversion fail", err)
	}
	return json
}
