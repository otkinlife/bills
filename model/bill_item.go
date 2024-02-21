package model

import "time"

type Bill struct {
	ID              string    `gorm:"type:varchar(128);not null"`
	ChargeTime      time.Time `gorm:"type:datetime;not null"`
	Value           float64   `gorm:"type:double;default:0;not null"`
	Channel         string    `gorm:"type:varchar(32);default:'';not null"`
	Goods           string    `gorm:"type:varchar(32);default:'';not null"`
	Tag             string    `gorm:"type:varchar(128);not null"`
	TransactionType string    `gorm:"type:varchar(32);default:'';not null"`
	SubType         string    `gorm:"type:varchar(32);default:'';not null"`
	Status          string    `gorm:"type:varchar(32);default:'';not null"`
	Name            string    `gorm:"type:varchar(32);default:'';not null"`
	Notes           string    `gorm:"type:text;not null"`
}

func (Bill) TableName() string {
	return "bill_data"
}
