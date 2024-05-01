package data

import "time"

type Exam struct {
	ID            int64		`json:"id"`
	Title         string	`json:"title"`
	Year		  int32 	`json:"year,omitempty"`
	Description   string	`json:"description"`
	NumberOfTests int32		`json:"number_of_tests"`
	CreatedAt     time.Time `json:"-"`
}