package database

import (
	"database/sql"

	"go-api-practice-5/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const createTablesSQL = `
CREATE TABLE IF NOT EXISTS merchandise (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	category VARCHAR(100),
	price INTEGER NOT NULL CHECK (price >= 0),
	description TEXT,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS characters (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	merchandise_id INTEGER NOT NULL REFERENCES merchandise(id) ON DELETE RESTRICT,
	intro TEXT,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
`

const seedMerchandiseSQL = `
DO $$
BEGIN
  IF (SELECT COUNT(*) FROM merchandise) = 0 THEN
    INSERT INTO merchandise (name, category, price, description) VALUES
      ('Hello Kitty 玩偶', '玩偶', 399, '經典款 25cm 絨毛玩偶'),
      ('美樂蒂 手帳本', '文具', 280, 'A6 尺寸，內附貼紙'),
      ('庫洛米 帆布包', '服飾', 520, '黑色帆布托特包');
  END IF;
END $$;
`

// Connect 建立 PostgreSQL 連線、建表，並塞入預設周邊資料
func Connect() error {
	var err error
	DB, err = sql.Open("postgres", config.DatabaseURL())
	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	if _, err = DB.Exec(createTablesSQL); err != nil {
		return err
	}
	_, err = DB.Exec(seedMerchandiseSQL)
	return err
}
