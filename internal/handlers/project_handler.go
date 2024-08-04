package handlers

import (
	"contact-center-system/internal/models"
	"contact-center-system/pkg"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"net/http"
)

type ProjectHandler struct {
	DB *bun.DB
}

func NewProjectHandler(db *bun.DB) *ProjectHandler {
	return &ProjectHandler{DB: db}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if _, err := h.DB.NewInsert().
		Model(&input).
		Exec(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
		"project": input,
	})
}

func (h *ProjectHandler) GetProject(c *gin.Context) {
	var projects []models.Project
	err := h.DB.NewSelect().
		Model(&projects).
		Relation("Operators").
		Order("name ASC").
		Scan(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, projects)
}

func (h *ProjectHandler) PutProject(c *gin.Context) {
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := pkg.ParseUUIDParam(c, "id")
	if err != nil {
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

	c.JSON(http.StatusOK, gin.H{"message": "Проект успешно обновлён"})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id, err := pkg.ParseUUIDParam(c, "id")
	if err != nil {
		return
	}
	err = pkg.DeleteByID(h.DB, &models.Project{}, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Проекта успешно удалён"})
}

func (h *ProjectHandler) DeleteOperatorProject(c *gin.Context) {
	operatorID, err := pkg.ParseUUIDParam(c, "operatorId")
	if err != nil {
		return
	}
	projectID, err := pkg.ParseUUIDParam(c, "id")
	if err != nil {
		return
	}
	result, err := h.DB.NewDelete().
		Model(&models.ProjectAssignment{}).
		Where("project_id = ? AND operator_id = ?", projectID, operatorID).
		Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = pkg.CheckUpdatedRows(result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": err})
}

func (h *ProjectHandler) AddOperatorProject(c *gin.Context) {
	operatorID, err := pkg.ParseUUIDParam(c, "operatorId")
	if err != nil {
		return
	}
	projectID, err := pkg.ParseUUIDParam(c, "id")
	if err != nil {
		return
	}
	var operatorDb models.Operator
	err = h.DB.NewSelect().
		Model(&operatorDb).
		Where("id = ?", operatorID).
		Scan(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	var projectDb models.Project
	err = h.DB.NewSelect().
		Model(&projectDb).
		Where("id = ?", projectID).
		Scan(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	assignment := &models.ProjectAssignment{
		ProjectID:  projectID,
		OperatorID: operatorID,
	}
	_, err = h.DB.NewInsert().
		Model(assignment).
		Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Оператор успешно добавлен к проекту",
	})

}
