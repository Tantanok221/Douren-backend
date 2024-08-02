package routes

import (
	"github.com/tantanok221/douren-backend/models"
	"net/http"

	"github.com/tantanok221/douren-backend/db"
	"github.com/tantanok221/douren-backend/handler"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	db := db.Init()
	Handler := &handler.ArtistHandler{
		DB: db,
	}
	mux.Handle("GET /artist", handler.GetArtist[models.PrimitiveArtist](Handler))
	mux.Handle("GET /event/{event}/artist", handler.HandleEventArtist(Handler))
	mux.Handle("GET /artist/{id}", Handler.GetArtistById())
	return mux
}
