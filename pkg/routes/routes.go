package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/quantinium03/conduit/pkg/handler"
)

func SetupRoutes(v1 chi.Router) {
	v1.Get("/health", handler.Health)

	// rawdog video transfers
	v1.Get("/{path}/direct", handler.DirectStream)

	//m3u8 file fetching
	v1.Get("/{path}/m3u8", handler.Getm3u8)

	// get media metadata
	v1.Get("/{path}/metadata", handler.GetMetaData)
}
