package models

import "math"

type Seat int

func (s Seat)SeatContains(a []Seat) bool {
	for _, n := range a {
		if s == n {
			return true
		}
	}

	return false
}

type Order struct {
	Dish string `json:"dish,omitempty" bson:"dish,omitempty"`
	Cost float64 `json:"cost,omitempty" bson:"cost,omitempty"`
}

// Get divided price on the number of participants
func SplitSharedOrderPrice(order Order, persons int) float64 {
	order.Cost = math.Round(order.Cost / float64(persons)*100)/100

	return order.Cost
}