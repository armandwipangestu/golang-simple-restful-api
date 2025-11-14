package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/armandwipangestu/golang-simple-restful-api/internal/models"
	"github.com/armandwipangestu/golang-simple-restful-api/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentHandler struct {
	Repo *repository.StudentRepo
}

func NewStudentHandler(r *repository.StudentRepo) *StudentHandler {
	return &StudentHandler{Repo: r}
}

func (h *StudentHandler) List(c *gin.Context) {
	students, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetById(c *gin.Context) {
	idParam := c.Param("id")

	// Convert string to uint
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
	}

	student, err := h.Repo.GetById(id)
	if err != nil {
		// If data not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Create(c *gin.Context) {
	var payload struct {
		Name	string	`json:"name" binding:"required"`
		Age		int		`json:"age" binding:"required"`
		Address struct {
			City	string	`json:"city"`
			Street	string	`json:"street"`
		} `json:"address"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		Name: payload.Name,
		Age: payload.Age,
		Address: models.Address{
			City: payload.Address.City,
			Street: payload.Address.Street,
		},
	}

	created, err := h.Repo.Create(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}