package entity

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	PaymentMethodName string `json:"status_method_name"` // ชื่อสถานะการชำระเงิน
	PaymentMethodNote string `json:"status_method_note"` // หมายเหตุสถานะการชำระเงิน

	Payments []Payment `gorm:"foreignKey:PaymentMethodID"` // Many-to-One relationship with Payment
}
