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
	var newMerchandise models.Merchandise
	if err := database.DB.QueryRow(query, input.Name, input.Category, input.Price, input.Description).Scan(&newMerchandise.ID, &newMerchandise.Name, &newMerchandise.Category, &newMerchandise.Price, &newMerchandise.Description, &newMerchandise.CreatedAt, &newMerchandise.UpdatedAt); err != nil {
		respondError(c, fmt.Errorf("新增周邊失敗：%w", err))
		return
	}

	c.JSON(http.StatusCreated, newMerchandise)
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

	query := "UPDATE merchandise SET name = $1, category = $2, price = $3, description = $4, updated_at = NOW() WHERE id = $5 RETURNING id, name, COALESCE(category, ''), price, COALESCE(description, ''), created_at, updated_at"

	var updateMerchandise models.Merchandise
	if err := database.DB.QueryRow(query, input.Name, input.Category, input.Price, input.Description, id).Scan(&updateMerchandise.ID, &updateMerchandise.Name, &updateMerchandise.Category, &updateMerchandise.Price, &updateMerchandise.Description, &updateMerchandise.CreatedAt, &updateMerchandise.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(c, ErrNotFound)
			return
		}
		respondError(c, fmt.Errorf("更新周邊失敗：%w", err))
		return
	}

	c.JSON(http.StatusOK, updateMerchandise)
}

// DeleteMerchandise 刪除周邊
func DeleteMerchandise(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	query := "DELETE FROM merchandise WHERE id = $1"

	result, err := database.DB.Exec(query, id)
	if err != nil {
		respondError(c, fmt.Errorf("刪除周邊失敗：%w", err))
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		respondError(c, fmt.Errorf("取得受影響行數失敗：%w", err))
		return
	}

	if rowsAffected == 0 {
		respondError(c, ErrNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "成功刪除"})
}
