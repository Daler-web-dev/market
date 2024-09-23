package models

import "time"

type Discount struct {
	Id              uint      `json: "id" gorm: "type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT" `
	Qty             int       `json: "qty"`
	Type            string    `json: "type"`
	Result          int       `json: "result"`
	ExpiredAt       int       `json: "expiredAt"`
	ExpiredAtFormat int       `json: "expiredAtFormat"`
	StringFormat    int       `json: "stringFormat"`
	CreatedAt       time.Time `json: "createdAt"`
	UpdatedAt       time.Time `json: "updatedAt"`
}
