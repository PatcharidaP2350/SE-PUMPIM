package entity

import (
	"gorm.io/gorm"
)

type StatusPayment struct {
	gorm.Model
	StatusPaymentName string `json:"status_payment_name"` // ชื่อสถานะการชำระเงิน
	StatusPaymentNote string `json:"status_payment_note"` // หมายเหตุสถานะการชำระเงิน

	Payments []Payment `gorm:"foreignKey:StatusPaymentID"` // Many-to-One relationship with Payment
}
