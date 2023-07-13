package model

import (
	"time"
)

type Coins struct {
	Id           int       `gorm:"primarykey" json:"id"`
	Name         string    `gorm:"type:varchar(50)" json:"name"`
	Abbreviation string    `gorm:"index" json:"abbreviation"`
	Symbol       string    `gorm:"type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci" json:"symbol"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Logs struct {
	Id        int       `gorm:"primarykey" json:"id"`
	Amount    float64   `json:"amount"`
	From      string    `gorm:"index" json:"from"`
	To        string    `gorm:"index" json:"to"`
	Rate      float64   `json:"rate"`
	CreatedAt time.Time `json:"createdAt"`
}
