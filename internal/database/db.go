package database

import (
	"database/sql"
	"fmt"
	"log"

	// Импортируем стандартный драйвер для pgx
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

// ConnectDB устанавливает соединение с базой данных и возвращает объект *bun.DB
func ConnectDB() *bun.DB {
	// Определяем строку подключения к базе данных
	dsn := "postgres://postgres:129056ss@localhost:5432/contact_center_db?sslmode=disable"
	// Открываем соединение с базой данных
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	// Проверяем соединение с базой данных
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Ошибка при проверке соединения с базой данных: %v", err)
	}
	// Создаем новый объект Bun DB с использованием PostgreSQL диалекта
	db := bun.NewDB(sqlDB, pgdialect.New())
	// Выводим сообщение об успешном подключении
	fmt.Println("Успешное подключение к базе данных.")
	// Возвращаем объект Bun DB
	return db
}
