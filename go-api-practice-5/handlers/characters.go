package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

	query := "SELECT * FROM character"





	respondError(c, fmt.Errorf("請實作：查詢 characters 並 LEFT JOIN merchandise 取得 merchandise_name，回傳列表"))
}

// GetCharacterByID 依 ID 取得單一角色
func GetCharacterByID(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	respondError(c, fmt.Errorf("請實作：依 id 查詢單一角色（id=%d）", id))
}

// CreateCharacter 新增角色
func CreateCharacter(c *gin.Context) {
	var input models.Character
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	respondError(c, fmt.Errorf("請實作：INSERT characters 並回傳（name=%s merchandise_id=%d）", input.Name, input.MerchandiseID))
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
	respondError(c, fmt.Errorf("請實作：UPDATE 角色並回傳（id=%d）", id))
}

// DeleteCharacter 刪除角色
func DeleteCharacter(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	respondError(c, fmt.Errorf("請實作：DELETE 角色（id=%d）", id))
}
