package models

import "time"

// Character 三麗鷗角色（關聯最愛周邊）
type Character struct {
	ID             int       `json:"id"`
	Name           string    `json:"name" binding:"required"`
	MerchandiseID  int       `json:"merchandise_id" binding:"required"`
	Intro          string    `json:"intro"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CharacterWithMerchandise 角色 + 周邊名稱（JOIN 查詢用）
type CharacterWithMerchandise struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	MerchandiseID    int       `json:"merchandise_id"`
	MerchandiseName  string    `json:"merchandise_name"`
	Intro            string    `json:"intro"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
