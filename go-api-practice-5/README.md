# Go API Practice 5 － 三麗鷗角色與周邊商品

Golang 後端 API 練習專案，主題為**三麗鷗角色**與**周邊商品**。  
角色可關聯一個「最愛周邊」；API 提供 characters 與 merchandise 的完整 CRUD（共 10 個路由）。

## 主題簡介

- **周邊商品（merchandise）**：三麗鷗周邊（玩偶、文具、服飾等）。可透過 API 查詢／新增／更新／刪除；啟動時若表為空會寫入 3 筆預設周邊。
- **角色（characters）**：三麗鷗角色，擁有一個最愛周邊，透過 `merchandise_id` 關聯。列表會 JOIN merchandise 帶出周邊名稱。

## 環境需求

- Go 1.21+
- PostgreSQL

## 設定

1. 建立 PostgreSQL 資料庫（例如名稱 `practice`）。
2. 複製並編輯 `.env`：
   ```bash
   cp .env.example .env
   ```
   - `PORT`：API 監聽 port（預設 `8080`）
   - `DATABASE_URL`：PostgreSQL 連線字串

## 執行

```bash
go mod tidy
go run main.go
```

伺服器預設在 `http://localhost:8080`。啟動時會自動建立 `merchandise`、`characters` 表，並在 merchandise 為空時寫入 3 筆預設周邊。

## Postman 測試

專案內含 Postman collection，可匯入後直接打 API：

- 檔案：`postman/go-api-practice-5.json`
- 在 Postman 選擇 **Import** → 選取上述 JSON 即可。
- 內含變數 `baseUrl`（預設 `http://localhost:8080`）、`characterId`、`merchandiseId`，可依需要修改。

## 專案結構

| 目錄／檔案 | 說明 |
|------------|------|
| `handlers/` | characters.go（角色 CRUD）、merchandise.go（周邊 CRUD）；GET 列表需實作 JOIN 帶出 merchandise_name |
| `models/` | Merchandise、Character、CharacterWithMerchandise（JOIN 用） |
| `database/` | 連線、建表、**merchandise 預設種子資料** |
| `routes/` | 10 個路由，皆已接上對應 handler；另有註解說明可選擴充路由 |

## API 路由（10 個）

| 方法 | 路徑 | 說明 |
|------|------|------|
| GET | `/api/characters` | 取得所有角色（JOIN merchandise 帶出 merchandise_name） |
| GET | `/api/characters/:id` | 依 ID 取得單一角色 |
| POST | `/api/characters` | 新增角色（body：`name`、`merchandise_id` 必填，`intro` 選填） |
| PUT | `/api/characters/:id` | 更新角色 |
| DELETE | `/api/characters/:id` | 刪除角色 |
| GET | `/api/merchandise` | 取得所有周邊 |
| GET | `/api/merchandise/:id` | 依 ID 取得單一周邊 |
| POST | `/api/merchandise` | 新增周邊（body：`name`、`price` 必填，`category`、`description` 選填） |
| PUT | `/api/merchandise/:id` | 更新周邊 |
| DELETE | `/api/merchandise/:id` | 刪除周邊 |

## 練習重點

- 在 `handlers/characters.go` 中，需實作 GET 列表（含 JOIN merchandise 取得 merchandise_name）；以及 GET :id、POST、PUT、DELETE 對 `characters` 的 DB 操作。
- 在 `handlers/merchandise.go` 中，需實作周邊的 CRUD（查詢列表、單筆、新增、更新、刪除）。
- `merchandise` 表由 `database.Connect()` 在啟動時若為空則塞入 3 筆預設周邊。
- 找不到資料時回傳 `ErrNotFound`，handler 會回 404。
