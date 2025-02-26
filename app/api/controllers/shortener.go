package controllers

import (
	"encoding/json"
	"net/http"
	"url-shortener/app/api/responses"
	"url-shortener/app/services"
)

type APIParams struct {
	URL       []string
	ShortKeys []string
}

func (p *APIParams) GetURL() []string {
	return p.URL
}

func (p *APIParams) GetShortKeys() []string {
	return p.ShortKeys
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	var urls []string
	if err := json.NewDecoder(r.Body).Decode(&urls); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	params := &APIParams{URL: urls, ShortKeys: nil}

	shortURLs, err := services.Shorten(params)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, shortURLs)
}
