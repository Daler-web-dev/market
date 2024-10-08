package models

import "time"

type Cashier struct {
	Id        uint      `json: "id"`
	Name      string    `json: "name"`
	Passcode  string    `json: "passcode"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
}
