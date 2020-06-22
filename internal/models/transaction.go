package models

import (
	"github.com/jinzhu/gorm"
)

type (
	Transaction struct {
		gorm.Model
		UserID        int   `json:"userId"`
		MerchantID    int   `json:"merchantId"`
		AmountInCents int64 `json:"amountInCents"`
		Timestamp     int64 `json:"timestamp"`
	}
)
