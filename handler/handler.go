package handler

import (
	"net/http"
	"strconv"

	"github.com/tantanok221/douren-backend/api"
	"github.com/tantanok221/douren-backend/internal/jsonlib"
	"github.com/uptrace/bun"
)

type ArtistHandler struct {
	DB *bun.DB
}

func (h ArtistHandler) GetAllArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := api.GetAllPrimitiveArtist(h.DB)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonlib.GetJson(data))
	}
}

func (h ArtistHandler) GetArtistById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if len(string(id)) == 0 || err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Missing Required Parameter"))
			return
		}

		data := api.GetPrimitiveArtistById(h.DB,id)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonlib.GetJson(data))
	}
}
