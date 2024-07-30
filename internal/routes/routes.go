package routes

import (
	"net/http"

	"github.com/tantanok221/douren-backend/db"
	"github.com/tantanok221/douren-backend/handler"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	db := db.Init()
	handler := &handler.ArtistHandler{
		DB: db,
	}
	mux.Handle("GET /artist", handler.GetAllArtist())
	mux.Handle("GET /artist/{id}", handler.GetArtistById())
	return mux
}
