package routes

import (
	"go-api-practice-5/handlers"

	"github.com/gin-gonic/gin"
)

// Setup 註冊 API 路由。以下依 handler 的 function 宣告大部分 route，
// 若有擴充需求可參考註解區塊。
func Setup(r *gin.Engine) {
	api := r.Group("/api")

	// --- 角色 (characters)：共 5 個 route，皆已接上 handler ---
	api.GET("/characters", handlers.GetCharacters)
	api.GET("/characters/:id", handlers.GetCharacterByID)
	api.POST("/characters", handlers.CreateCharacter)
	api.PUT("/characters/:id", handlers.UpdateCharacter)
	api.DELETE("/characters/:id", handlers.DeleteCharacter)

	// --- 周邊商品 (merchandise)：共 5 個 route，皆已接上 handler ---
	api.GET("/merchandise", handlers.GetMerchandise)
	api.GET("/merchandise/:id", handlers.GetMerchandiseByID)
	api.POST("/merchandise", handlers.CreateMerchandise)
	api.PUT("/merchandise/:id", handlers.UpdateMerchandise)
	api.DELETE("/merchandise/:id", handlers.DeleteMerchandise)

	// --- 以下為「剩下要宣告的 route」說明（可選擴充）---
	// 若需要依分類查周邊：GET /api/merchandise/category/:category -> 需新增 handler 如 GetMerchandiseByCategory。
	// 若需要依角色查周邊：GET /api/characters/:id/merchandise -> 回傳該角色關聯的周邊，需新增 handler。
}
