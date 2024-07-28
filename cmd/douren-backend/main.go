package main

import (
	"net/http"

	"github.com/tantanok221/douren-backend/db"
	"github.com/tantanok221/douren-backend/handler"
)

func main() {
	mux := http.NewServeMux()
	print("Server Running")
	handler := &handler.ArtistHandler{
		DB:  db.Init(),
	}
	mux.Handle("GET /artist", handler.GetAllArtist())
	mux.Handle("GET /artist/{id}", handler.GetArtistById())
	http.ListenAndServe(":3000",mux)
}
