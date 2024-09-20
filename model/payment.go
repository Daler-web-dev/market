package model

import "time"

type Payment struct {
	Id            uint      `json: "id" gorm: "type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT" `
	Name          string    `json: "name"`
	Type          string    `json: "type"`
	PaymentTypeId int       `json: "payment_type_id"`
	Logo          string    `json: "logo"`
	CreatedAt     time.Time `json: "createdAt"`
	UpdatedAt     time.Time `json: "updatedAt"`
}

type PaymentType struct {
	Id        uint      `json: "id" gorm: "type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT" `
	Name      string    `json: "name"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
}
