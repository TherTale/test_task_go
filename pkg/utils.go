package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"net/http"
)

// Contains принимает срез любого типа T и значение того же типа для поиска.
func Contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// ParseUUIDParam извлекает и парсит UUID из параметра запроса.
func ParseUUIDParam(c *gin.Context, paramName string) (uuid.UUID, error) {
	param := c.Param(paramName)
	id, err := uuid.Parse(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID " + paramName})
		return uuid.UUID{}, err
	}
	return id, nil
}

func DeleteByID(db bun.IDB, model interface{}, id uuid.UUID) error {
	result, err := db.NewDelete().
		Model(model).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось удалить запись")
	}
	err = CheckUpdatedRows(result)
	if err != nil {
		return err
	}
	return nil
}

func CheckUpdatedRows(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("jшибка при проверке удаления")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("данные не найдены")
	}
	return nil
}
