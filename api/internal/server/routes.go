package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.opencensus.io/plugin/ochttp"
)

func Routes() (http.Handler, error) {
	r := chi.NewRouter()
	r.Get("/api/getExample", handleGetExample)
	r.Post("/api/postExample", handlePostExample)
	return &ochttp.Handler{Handler: r}, nil
}
