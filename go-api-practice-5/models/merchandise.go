package models

import "time"

// Merchandise 三麗鷗周邊商品
type Merchandise struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category"`     // 如：玩偶、文具、服飾
	Price       int       `json:"price" binding:"required"` // 單位：元
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
