package repository

import (
	"database/sql"

	"github.com/santtisosa/trackr/backend/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetCategoriesByUserID(userID string) ([]model.Category, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, name, icon, color, is_default, created_at, updated_at
		FROM categories
		WHERE user_id = $1
		ORDER BY name ASC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.UserID, &c.Name, &c.Icon, &c.Color, &c.IsDefault, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, rows.Err()
}
