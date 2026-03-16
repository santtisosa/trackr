package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santtisosa/trackr/backend/internal/model"
	"github.com/santtisosa/trackr/backend/internal/repository"
)

type ExpenseHandler struct {
	repo *repository.ExpenseRepository
}

func NewExpenseHandler(repo *repository.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repo}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	userID := c.GetString("user_id")

	var input model.Expense
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = userID

	expense, err := h.repo.CreateExpense(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create expense"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	userID := c.GetString("user_id")

	expenses, err := h.repo.GetExpensesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch expenses"})
		return
	}

	if expenses == nil {
		expenses = []model.Expense{}
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *ExpenseHandler) GetExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	id := c.Param("id")

	expense, err := h.repo.GetExpenseByID(id, userID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch expense"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	id := c.Param("id")

	var input model.Expense
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	input.UserID = userID

	expense, err := h.repo.UpdateExpense(input)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update expense"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	userID := c.GetString("user_id")
	id := c.Param("id")

	err := h.repo.DeleteExpense(id, userID)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete expense"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
