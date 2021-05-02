package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
)

type Table struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Active bool               `json:"active,omitempty" bson:"active,omitempty"`
	Table  int64              `json:"table,omitempty" bson:"table,omitempty"`
	Guests []Guest            `json:"guests,omitempty" bson:"guests,omitempty"`
}

type TableOrder []struct {
	Seats  []Seat  `json:"seats,omitempty" bson:"seats,omitempty"`
	Shared bool    `json:"shared,omitempty" bson:"shared,omitempty"`
	Orders []Order `json:"orders,omitempty" bson:"orders,omitempty"`
}

// Prepare bill with list and amount for seats
func (t *Table) ProcessBills(bsr *BillsRequest) []*Bill {
	bills := make([]*Bill, 0)
	for _, s := range bsr.Seats {

		bill := new(Bill)
		for _, guest := range t.Guests {

			if guest.Seat.SeatContains(s) {
				for _, order := range guest.Orders {
					bill.Orders = append(bill.Orders, order)
					bill.Amount += order.Cost
				}
			}

		}

		bill.Amount = math.Round(bill.Amount*100)/100
		bill.Seats = s
		bills = append(bills, bill)
	}

	return bills
}

// Split all shared orders price by the number of participants
func (to *TableOrder) ProcessSharedOrder() {
	for _, order := range *to {

		//if order.Shared == true {
		if order.Shared {
			for j, odr := range order.Orders {
				order.Orders[j].Cost = SplitSharedOrderPrice(odr, len(order.Seats))
			}
		}

	}
}

// Add order to the correct guest (seat) on the table
func (to *TableOrder) AddOrderToGuest(guest *Guest) {
	for _, order := range *to {

		for _, seat := range order.Seats {
			if seat == guest.Seat {
				guest.Orders = append(guest.Orders, order.Orders...)
			}
		}

	}
}
