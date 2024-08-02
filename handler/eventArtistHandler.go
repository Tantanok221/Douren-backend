package handler

import (
	"github.com/tantanok221/douren-backend/api"
	"github.com/tantanok221/douren-backend/internal/helper"
	"github.com/tantanok221/douren-backend/models"
	"net/http"
	"strconv"
)

func HandleEventArtist(h *ArtistHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, page, err := parsePaginationParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		eventId, err := strconv.Atoi(r.PathValue("event"))
		if err != nil || string(eventId) == "" {
			http.Error(w, "Missing Required Parameter", http.StatusUnprocessableEntity)
			return
		}
		options := &api.DBOptions{
			Limit: limit,
			Page:  page,
			DB:    h.DB,
		}
		dataResponse := api.EventFetch(options, eventId)
		paginationResponse := api.GetPagination[models.EventArtist](options)
		jsonResponse := APIWrapper[[]models.EventArtist]{
			Data:       dataResponse,
			Pagination: paginationResponse,
		}

		helper.WriteJSON(w, jsonResponse)
	}
}
