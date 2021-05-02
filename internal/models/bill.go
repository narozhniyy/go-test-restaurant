package models

type Bill struct {
	Seats  []Seat  `json:"seats,omitempty"`
	Orders []Order `json:"orders,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type BillsRequest struct {
	Seats [][]Seat `json:"bills"`
}
