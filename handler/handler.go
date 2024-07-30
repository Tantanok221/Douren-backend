package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/tantanok221/douren-backend/api"
	"github.com/tantanok221/douren-backend/internal/helper"
	"github.com/uptrace/bun"
)

type ArtistHandler struct {
	DB *bun.DB
}

func (h ArtistHandler) GetAllArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, page, err := parseAllArtistParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		options := &api.GetAllPrimitiveArtistOptions{
			DB:    h.DB,
			Limit: limit,
			Page:  page,
		}
		data := options.GetAllPrimitiveArtist()

		helper.WriteJSON(w, data)
	}
}

func parseAllArtistParams(r *http.Request) (limit, page int, err error) {
	limit, err = helper.HandleParam(r.URL.Query().Get("limit"), 20)
	if err != nil {
		return 0, 0, errors.New("limit parameter is not a number")
	}

	page, err = helper.HandleParam(r.URL.Query().Get("page"), 0)
	if err != nil {
		return 0, 0, errors.New("page parameter is not a number")
	}

	return limit, page, nil
}

func (h ArtistHandler) GetArtistById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if string(id) == "" || err != nil {
			http.Error(w,"Missing Required Parameter", http.StatusUnprocessableEntity)
			return
		}

		data := api.GetPrimitiveArtistById(h.DB, id)
		helper.WriteJSON(w,data)
	}
}
