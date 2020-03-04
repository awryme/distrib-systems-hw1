package handler_inc

import (
	"go.uber.org/zap"
	"incrementer/application"
	"incrementer/log"
	"net/http"
)

func GetHandlerInc(app *application.Application) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := log.FromCtx(ctx)

		reqObj, err := parseRequestObj(r.Body)
		if err != nil {
			logErr(logger, "failed to get request", err)
			sendErr(w, http.StatusBadRequest, err)
			return
		}

		res, err := app.IncrementV1(ctx, reqObj.Query)
		if err != nil {
			logErr(logger, "failed to increment", err)
			sendErr(w, http.StatusInternalServerError, err)
			return
		}

		logOk(logger, reqObj.Query, res)
		sendOk(w, reqObj.Query, res)
	}
}

func sendErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	renderErr(w, err)
}

func logErr(logger *zap.Logger, msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

func logOk(logger *zap.Logger, query, resp int) {
	logger.Info("incremented",
		zap.Int("query", query),
		zap.Int("resp", resp),
	)
}

func sendOk(w http.ResponseWriter, query, resp int) {
	w.WriteHeader(http.StatusOK)
	res := okResponse{
		Query: query,
		Resp:  resp,
	}
	renderJson(w, res)
}
