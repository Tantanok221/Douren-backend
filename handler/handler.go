package handler

import (
	"net/http"
	"github.com/tantanok221/douren-backend/data/fake"
)

func GetAllArtist() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json,_ := fake.ReturnJson()
		w.Write(json)
	}
}
