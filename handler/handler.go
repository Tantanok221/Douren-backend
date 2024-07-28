package handler

import (
	"net/http"

	"github.com/tantanok221/douren-backend/api"
	"github.com/uptrace/bun"
)

type ArtistHandler struct {
	DB *bun.DB
}

func (h ArtistHandler) GetAllArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json := api.GetAllPrimitiveArtist(h.DB)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

func (h ArtistHandler) GetArtistById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json := api.GetAllPrimitiveArtist(h.DB)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}
