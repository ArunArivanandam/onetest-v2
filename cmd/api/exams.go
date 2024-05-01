package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arunarivanandam/onetest-v2/internal/data"
)

func (app *application) createExamHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title         string `json:"title"`
		Year		  int32  `json:"year"`
		Description   string `json:"description"`
		NumberOfTests int32  `json:"number_of_tests"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v", input)
}

func (app *application) showExamHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	exam := data.Exam{
		ID:            id,
		Title:         "TNPSC GROUP 4 EXAM",
		Year:		  	2024,
		Description:   "2024 TNPSC GROUP 4 EXAM MOCK TESTS",
		NumberOfTests: 20,
		CreatedAt:     time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"exam": exam}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
