package server

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"incrementer/application"
	"incrementer/server/handler-inc"
	"net/http"
)


func NewServer(logger *zap.Logger, app *application.Application) http.Handler{
	router := chi.NewRouter()
	register(router, logger, app)
	return router
}

func register(r chi.Router, logger *zap.Logger, app *application.Application) {
	r.Use(middlewareUseLogger(logger))
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(middlewareSetApiVersion("v1"))
			r.Post("/inc", handler_inc.GetHandlerInc(app))
		})
		r.Route("/v2", func(r chi.Router) {
			r.Use(middlewareSetApiVersion("v2"))
			r.Post("/inc", handler_inc.GetHandlerInc(app))
		})
	})
}