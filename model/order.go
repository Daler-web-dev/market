package model

import "time"

type Order struct {
	Id            uint      `json: "id" gorm: "type: INT(10) UNSIGNED NOT NULL AUTO_INCREMENT" `
	CashierID     int       `json: "cashierId"`
	PaymentTypeId int       `json: "paymentTypeId"`
	TotalPrice    int       `json: "totalPrice"`
	TotalPaid     int       `json: "totalPaid"`
	TotalReturn   int       `json: "totalReturn"`
	ReceiptId     string    `json: "receiptId"`
	IsDownload    int       `json: "is_download"`
	ProductId     string    `json: "productId"`
	Quantities    string    `json: "quantities"`
	CreatedAt     time.Time `json: "created_at"`
	UpdatedAt     time.Time `json: "updated_at"`
}
