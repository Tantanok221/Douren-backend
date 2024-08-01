package handler

import (
	"errors"
	"github.com/tantanok221/douren-backend/models"
	"net/http"
	"strconv"

	"github.com/tantanok221/douren-backend/api"
	"github.com/tantanok221/douren-backend/internal/helper"
	"github.com/uptrace/bun"
)

type ArtistHandler struct {
	DB *bun.DB
}

type APIWrapper[T any] struct {
	Data       T                 `json:"data"`
	Pagination models.Pagination `json:"pagination"`
}

func (h ArtistHandler) GetAllArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, page, err := parseAllArtistParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		options := &api.DBOptions{
			DB:    h.DB,
			Limit: limit,
			Page:  page,
		}
		dataResponse := options.GetAllPrimitiveArtist()
		paginationResponse := api.GetPagination[models.PrimitiveArtist](options)
		jsonResponse := APIWrapper[[]models.PrimitiveArtist]{
			Data:       dataResponse,
			Pagination: paginationResponse,
		}

		helper.WriteJSON(w, jsonResponse)
	}
}

func parseAllArtistParams(r *http.Request) (limit, page int, err error) {
	limit, err = helper.HandleParam(r.URL.Query().Get("limit"), 20)
	if err != nil {
		return 0, 0, errors.New("limit parameter is not a number")
	}

	page, err = helper.HandleParam(r.URL.Query().Get("page"), 1)
	if err != nil {
		return 0, 0, errors.New("page parameter is not a number")
	}

	return limit, page, nil
}

func (h ArtistHandler) GetArtistById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if string(id) == "" || err != nil {
			http.Error(w, "Missing Required Parameter", http.StatusUnprocessableEntity)
			return
		}

		data := api.GetPrimitiveArtistById(h.DB, id)
		helper.WriteJSON(w, data)
	}
}
