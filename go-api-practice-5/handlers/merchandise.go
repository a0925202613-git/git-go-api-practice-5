package handlers

import (
	"database/sql" //為了使用 sql.ErrNoRows
	"errors"
	"fmt"
	"net/http"

	"go-api-practice-5/database" //為了呼叫 database.DB
	"go-api-practice-5/models"

	"github.com/gin-gonic/gin"
)

// GetMerchandise 取得所有周邊商品
func GetMerchandise(c *gin.Context) {

	query := "SELECT id, name, COALESCE(category, ''), price, COALESCE(description, ''), created_at, updated_at FROM merchandise"

	rows, err := database.DB.Query(query)
	if err != nil {
		respondError(c, fmt.Errorf("取得所有周邊商品失敗:%w", err))
		return
	}
	defer rows.Close()

	var items []models.Merchandise
	for rows.Next() {
		var item models.Merchandise
		if err := rows.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.Description, &item.CreatedAt, &item.UpdatedAt); err != nil {
			respondError(c, fmt.Errorf("取得失敗:%w", err))
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

// GetMerchandiseByID 依 ID 取得單一周邊
func GetMerchandiseByID(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	query := "SELECT id, name, COALESCE(category, ''), price, COALESCE(description, ''), created_at, updated_at FROM merchandise WHERE id =$1"

	var item models.Merchandise
	if err := database.DB.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.Description, &item.CreatedAt, &item.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(c, ErrNotFound)
			return
		}
		respondError(c, fmt.Errorf("取得單一失敗:%w", err))
		return
	}

	c.JSON(http.StatusOK, item)
}

// CreateMerchandise 新增周邊（body：name、price 必填，category、description 選填）
func CreateMerchandise(c *gin.Context) {
	var input models.Merchandise
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

query := `
    INSERT INTO merchandise (name, category, price, description) 
    VALUES ($1, $2, $3, $4) 
    RETURNING id, name, COALESCE(category, ''), price, COALESCE(description, ''), created_at, updated_at
`
var newMerchandi models.Merchandise
if err := database.DB.QueryRow(query, input.ID)



	respondError(c, fmt.Errorf("請實作：INSERT merchandise 並回傳（name=%s price=%d）", input.Name, input.Price))
}

// UpdateMerchandise 更新周邊
func UpdateMerchandise(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var input models.Merchandise
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	respondError(c, fmt.Errorf("請實作：UPDATE 周邊並回傳（id=%d）", id))
}

// DeleteMerchandise 刪除周邊
func DeleteMerchandise(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	respondError(c, fmt.Errorf("請實作：DELETE 周邊（id=%d）", id))
}
