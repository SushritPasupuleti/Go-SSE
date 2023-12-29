package models

type Stock struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
	Value  int    `json:"value"`
	Timestamp string `json:"timestamp"`
}
