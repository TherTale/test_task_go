package handlers

import (
	"contact-center-system/internal/models"
	"contact-center-system/internal/services"
	"contact-center-system/pkg"
	"context"
	"github.com/gin-gonic/gin"
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
	err := h.DB.NewSelect().
		Model(&operators).
		Relation("Projects").
		Order("first_name ASC").
		Scan(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения оператора"})
		return
	}
	c.JSON(http.StatusOK, operators)
}

func (h *OperatorHandler) CreateOperators(c *gin.Context) {
	var input models.Operator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	passLength := 10
	generatedPassword := services.GeneratePassword(passLength)
	hashedPassword, err := services.HashPassword(generatedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	input.Password = hashedPassword
	_, err = h.DB.NewInsert().Model(&input).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Operator created successfully",
		"password": generatedPassword,
	})
}

func (h *OperatorHandler) PutOperators(c *gin.Context) {
	var input models.Operator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := pkg.ParseUUIDParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.DB.NewUpdate().
		Model(&input).
		Column(input.GetUpdateFields()...).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = pkg.CheckUpdatedRows(result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Оператор успешно обновлён"})
}

func (h *OperatorHandler) DeleteOperators(c *gin.Context) {
	id, err := pkg.ParseUUIDParam(c, "id")
	if err != nil {
		return
	}
	err = pkg.DeleteByID(h.DB, &models.Operator{}, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Оператор успешно удалён"})
}
