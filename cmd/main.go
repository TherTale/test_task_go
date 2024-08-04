package main

import (
	"contact-center-system/internal/database"
	"contact-center-system/internal/handlers"
	"contact-center-system/internal/models"
	routes "contact-center-system/internal/router"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// migrateDatabase выполняет миграции для базы данных
func migrateDatabase(dsn string) {
	// Открываем соединение с базой данных
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	// Создаем экземпляр драйвера для postgres
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Ошибка при создании драйвера миграции: %v", err)
	}

	// Указываем путь к миграциям и источник
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Ошибка при создании миграции: %v", err)
	}

	// Выполняем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Ошибка при выполнении миграции: %v", err)
	}
	log.Println("Миграции успешно выполнены")
}

func main() {
	// Определяем строку подключения к базе данных
	dsn := "postgres://postgres:129056ss@localhost:5432/contact_center_db?sslmode=disable"

	// Выполняем миграции перед подключением
	migrateDatabase(dsn)

	// Подключаемся к базе данных
	db := database.ConnectDB()
	db.RegisterModel((*models.ProjectAssignment)(nil))
	defer db.Close()

	operatorHandler := handlers.NewOperatorHandler(db)
	projectHandler := handlers.NewProjectHandler(db)

	// Инициализируем роутер Gin
	router := gin.Default()

	// Регистрируем маршруты
	routes.RegisterOperatorRoutes(router, operatorHandler)
	routes.RegisterProjectRoutes(router, projectHandler)

	// Запускаем сервер
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
