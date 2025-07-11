package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DB_DSN не указан в .env")
	}

	// + база
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии подключения: %w", err)
	}

	// чекаем соединение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", err)
	}

	// + таблицы если их нет
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("ошибка при создании таблиц: %w", err)
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	// + таблица с ключами
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS key_list (
			id SERIAL PRIMARY KEY,
			key_name TEXT UNIQUE NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("key_list: %w", err)
	}

	// + таблица с результатами
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS json_data (
			id SERIAL PRIMARY KEY,
			key_name TEXT NOT NULL,
			value TEXT
		);
	`)
	if err != nil {
		return fmt.Errorf("json_data: %w", err)
	}

	return nil
}

// возвращаем список нужных ключей из таблицы key_list
func GetWantedKeys(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`SELECT key_name FROM key_list`)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении ключей: %w", err)
	}
	defer rows.Close()

	var keys []string
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			return nil, fmt.Errorf("ошибка при чтении ключа: %w", err)
		}
		keys = append(keys, key)
	}

	return keys, nil
}

// вставляем отфильтрованные данные в json_data
func InsertFilteredData(db *sql.DB, key, value string) error {
	_, err := db.Exec(`INSERT INTO json_data (key_name, value) VALUES ($1, $2)`, key, value)
	if err != nil {
		return fmt.Errorf("ошибка вставки данных (%s): %w", key, err)
	}
	return nil
}
