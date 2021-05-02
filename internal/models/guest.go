package models

type Guest struct {
	Name   string  `json:"name,omitempty" bson:"name,omitempty"`
	Seat   Seat    `json:"seat,omitempty" bson:"seat,omitempty"`
	Orders []Order `json:"orders,omitempty" bson:"orders,omitempty"`
}
