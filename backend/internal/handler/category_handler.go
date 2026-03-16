package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santtisosa/trackr/backend/internal/model"
	"github.com/santtisosa/trackr/backend/internal/repository"
)

type CategoryHandler struct {
	repo *repository.CategoryRepository
}

func NewCategoryHandler(repo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	userID := c.GetString("user_id")

	categories, err := h.repo.GetCategoriesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch categories"})
		return
	}

	if categories == nil {
		categories = []model.Category{}
	}

	c.JSON(http.StatusOK, categories)
}
