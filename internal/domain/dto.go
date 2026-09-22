package domain

import "time"

type GameResult struct {
	Date     time.Time `json:"date"`
	Result   string    `json:"result"`
	Attempts int       `json:"attempts"`
}
