package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go-api-practice-5/database"
	"go-api-practice-5/models"

	"github.com/gin-gonic/gin"
)

var ErrNotFound = errors.New("not found")

func respondError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	if errors.Is(err, ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func parseID(c *gin.Context, param string) (int, bool) {
	id, err := strconv.Atoi(c.Param(param))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}

// GetCharacters 取得所有角色（需 JOIN merchandise 帶出周邊名稱，回傳 []CharacterWithMerchandise）
func GetCharacters(c *gin.Context) {

	query := `
		SELECT
			c.id,
			c.name,
			c.merchandise_id,
			COALESCE(m.name, '') AS merchandise_name,
			COALESCE(c.intro, ''),
			c.created_at,
			c.updated_at
		FROM characters c
		LEFT JOIN merchandise m ON c.merchandise_id = m.id
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		respondError(c, fmt.Errorf("取得所有角色失敗：%w", err))
		return
	}

	defer rows.Close()

	var chars []models.CharacterWithMerchandise

	for rows.Next() {
		var char models.CharacterWithMerchandise
		if err := rows.Scan(
			&char.ID,
			&char.Name,
			&char.MerchandiseID,
			&char.MerchandiseName,
			&char.Intro,
			&char.CreatedAt,
			&char.UpdatedAt,
		); err != nil {
			respondError(c, fmt.Errorf("讀取角色資料失敗：%w", err))
			return
		}

		chars = append(chars, char)
	}

	c.JSON(http.StatusOK, chars)
}

// GetCharacterByID 依 ID 取得單一角色
func GetCharacterByID(c *gin.Context) {
	// 驗證 ID 是否有效
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	// 使用id查詢角色資料，並 JOIN merchandise 帶出周邊名稱
	query := `
		SELECT
			c.id,
			c.name,
			c.merchandise_id,
			COALESCE(m.name, '') AS merchandise_name,
			COALESCE(c.intro, ''),
			c.created_at,
			c.updated_at
		FROM characters c
		LEFT JOIN merchandise m ON c.merchandise_id = m.id
		WHERE c.id = $1
	`

	// 準備盒子接收查詢結果
	var item models.CharacterWithMerchandise

	// 執行查詢，並將結果掃描到盒子裡
	if err := database.DB.QueryRow(query, id).Scan(
		&item.ID,
		&item.Name,
		&item.MerchandiseID,
		&item.MerchandiseName,
		&item.Intro,
		&item.CreatedAt,
		&item.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(c, ErrNotFound)
			return
		}
		respondError(c, fmt.Errorf("查詢角色失敗：%w", err))
		return
	}

	// 回傳查詢結果
	c.JSON(http.StatusOK, item)
}

// CreateCharacter 新增角色
func CreateCharacter(c *gin.Context) {
	var input models.Character
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 實作：INSERT characters 並回傳（name, merchandise_id 從 input 帶入）
	query := `
		INSERT INTO characters (name, merchandise_id, intro, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, name, merchandise_id, COALESCE(intro, ''), created_at, updated_at
	`

	// 準備盒子接收查詢結果
	var newchar models.CharacterWithMerchandise

	// 執行查詢，並將結果掃描到盒子裡
	if err := database.DB.QueryRow(query, input.Name, input.MerchandiseID, input.Intro).Scan(
		&newchar.ID,
		&newchar.Name,
		&newchar.MerchandiseID,
		&newchar.Intro,
		&newchar.CreatedAt,
		&newchar.UpdatedAt,
	); err != nil {
		respondError(c, fmt.Errorf("新增角色失敗：%w", err))
		return
	}

	c.JSON(http.StatusOK, newchar)
}

// UpdateCharacter 更新角色
func UpdateCharacter(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var input models.Character
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name == "" || input.MerchandiseID == 0 {
		// 查詢原本的角色資料，確保角色存在，並取得現有的 name 和 merchandise_id
		var existing models.Character
		err := database.DB.QueryRow(`SELECT name, merchandise_id FROM characters WHERE id = $1`, id).Scan(&existing.Name, &existing.MerchandiseID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				respondError(c, ErrNotFound)
				return
			}
			respondError(c, fmt.Errorf("查詢角色失敗：%w", err))
			return
		}
		input.Name = existing.Name
		input.MerchandiseID = existing.MerchandiseID
	}

	// 實作：UPDATE characters 並回傳（name, merchandise_id 從 input 帶入）
	query := `
		UPDATE characters
		SET name = $1, merchandise_id = $2, intro = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING id, name, merchandise_id, COALESCE(intro, ''), created_at, updated_at
	`

	// 準備盒子接收查詢結果
	var updatedchar models.CharacterWithMerchandise

	// 執行查詢，並將結果掃描到盒子裡
	if err := database.DB.QueryRow(query, input.Name, input.MerchandiseID, input.Intro, id).Scan(
		&updatedchar.ID,
		&updatedchar.Name,
		&updatedchar.MerchandiseID,
		&updatedchar.MerchandiseName,
		&updatedchar.Intro,
		&updatedchar.CreatedAt,
		&updatedchar.UpdatedAt,
	); err != nil {
		respondError(c, fmt.Errorf("更新角色失敗：%w", err))
		return
	}

	c.JSON(http.StatusOK, updatedchar)
}

// DeleteCharacter 刪除角色
func DeleteCharacter(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}

	// 實作：DELETE characters（依 id 刪除）
	query := `DELETE FROM characters WHERE id = $1`

	// 執行刪除(不需要回傳資料，用 Exec 就好)
	result, err := database.DB.Exec(query, id)
	if err != nil {
		respondError(c, fmt.Errorf("刪除角色失敗：%w", err))
		return
	}

	// 檢查是否有資料被刪除
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		respondError(c, fmt.Errorf("無法取得刪除結果：%w", err))
		return
	}

	// 如果沒有資料被刪除，代表找不到該角色
	if rowsAffected == 0 {
		respondError(c, ErrNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "角色已刪除"})

}
