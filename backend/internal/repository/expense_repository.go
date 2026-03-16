package repository

import (
	"database/sql"

	"github.com/santtisosa/trackr/backend/internal/model"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) CreateExpense(expense model.Expense) (model.Expense, error) {
	var e model.Expense
	err := r.db.QueryRow(`
		INSERT INTO expenses (user_id, category_id, amount, description, date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, category_id, amount, description, date, created_at, updated_at
	`, expense.UserID, expense.CategoryID, expense.Amount, expense.Description, expense.Date).
		Scan(&e.ID, &e.UserID, &e.CategoryID, &e.Amount, &e.Description, &e.Date, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (r *ExpenseRepository) GetExpensesByUserID(userID string) ([]model.Expense, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, category_id, amount, description, date, created_at, updated_at
		FROM expenses
		WHERE user_id = $1
		ORDER BY date DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []model.Expense
	for rows.Next() {
		var e model.Expense
		if err := rows.Scan(&e.ID, &e.UserID, &e.CategoryID, &e.Amount, &e.Description, &e.Date, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}
	return expenses, rows.Err()
}

func (r *ExpenseRepository) GetExpenseByID(id string, userID string) (model.Expense, error) {
	var e model.Expense
	err := r.db.QueryRow(`
		SELECT id, user_id, category_id, amount, description, date, created_at, updated_at
		FROM expenses
		WHERE id = $1 AND user_id = $2
	`, id, userID).
		Scan(&e.ID, &e.UserID, &e.CategoryID, &e.Amount, &e.Description, &e.Date, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (r *ExpenseRepository) UpdateExpense(expense model.Expense) (model.Expense, error) {
	var e model.Expense
	err := r.db.QueryRow(`
		UPDATE expenses
		SET category_id = $1, amount = $2, description = $3, date = $4, updated_at = NOW()
		WHERE id = $5 AND user_id = $6
		RETURNING id, user_id, category_id, amount, description, date, created_at, updated_at
	`, expense.CategoryID, expense.Amount, expense.Description, expense.Date, expense.ID, expense.UserID).
		Scan(&e.ID, &e.UserID, &e.CategoryID, &e.Amount, &e.Description, &e.Date, &e.CreatedAt, &e.UpdatedAt)
	return e, err
}

func (r *ExpenseRepository) DeleteExpense(id string, userID string) error {
	result, err := r.db.Exec(`
		DELETE FROM expenses WHERE id = $1 AND user_id = $2
	`, id, userID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
