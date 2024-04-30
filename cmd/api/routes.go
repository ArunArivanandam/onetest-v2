package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/exams", app.createExamHandler)
	router.HandlerFunc(http.MethodGet, "/v1/exams/:id", app.showExamHandler)

	return router
	

}