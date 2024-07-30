package jsonlib

import (
	"encoding/json"
	"log"
)

func GetJson[T any](result T) []byte {
	json, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Json Conversion fail", err)
	}
	return json
}
