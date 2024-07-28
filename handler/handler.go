package handler

import (
	"net/http"

	"github.com/tantanok221/douren-backend/api"
)

func HandleArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var json []byte
		if r.Method == http.MethodGet {
			json = api.GetAllPrimitiveArtist()
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
