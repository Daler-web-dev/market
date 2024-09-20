package model

import "time"

type Category struct {
	Id        uint      `json: "id" gorm: "type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT" `
	Name      string    `json: "name"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
}
