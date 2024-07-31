package handlers

import (
	"contact-center-system/internal/models"
	"contact-center-system/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"net/http"
	_ "strings"
)

type OperatorHandler struct {
	DB *bun.DB
}

func NewOperatorHandler(db *bun.DB) *OperatorHandler {
	return &OperatorHandler{DB: db}
}

func (h *OperatorHandler) GetOperators(c *gin.Context) {
	var operators []models.Operator
	err := h.DB.NewSelect().Model(&operators).Order("first_name ASC").Scan(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch operators"})
		return
	}
	c.JSON(http.StatusOK, operators)
}

func (h *OperatorHandler) CreateOperators(c *gin.Context) {
	var input models.Operator

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if err := services.ValidateOperator(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	generatedPassword := services.GeneratePassword(10) // Длина пароля 10 символов
	hashedPassword, err := services.HashPassword(generatedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	input.Password = hashedPassword
	_, err = h.DB.NewInsert().Model(&input).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create operator"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Operator created successfully",
		"password": generatedPassword, // Отправляем сгенерированный пароль пользователю
	})

}

func (h *OperatorHandler) PutOperators(c *gin.Context) {
	var input models.Operator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные входные данные"})
		return
	}
	if err := services.ValidateOperator(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operatorID := c.Param("id")
	id, err := uuid.Parse(operatorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID оператора"})
		return
	}

	var operator models.Operator
	err = h.DB.NewSelect().Model(&operator).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Оператор не найден"})
		return
	}

	operator.City = input.City
	operator.PhoneNumber = input.PhoneNumber
	operator.FirstName = input.FirstName
	operator.LastName = input.LastName
	operator.Email = input.Email
	operator.MiddleName = input.MiddleName

	_, err = h.DB.NewUpdate().Model(&operator).Where("id = ?", operator.ID).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить оператора"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Оператор успешно обновлён"})
}

func (h *OperatorHandler) DeleteOperators(c *gin.Context) {
	operatorID := c.Param("id")
	id, err := uuid.Parse(operatorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID оператора"})
		return
	}

	result, err := h.DB.NewDelete().Model(&models.Operator{}).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить оператора"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке удаления оператора"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Оператор не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Оператор успешно удалён"})
}
