package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Supplier  string    `json:"supplier"`
	Amount    int       `json:"amount"`
	CreatedOn time.Time `json:"createdon"`
	ChangedOn time.Time `json:"changedon"`

	
}

