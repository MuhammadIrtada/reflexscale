package entity

import "time"

type (
	HasilTest struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		UserID    uint      `json:"user_id" gorm:"index;not null"`
		Poin      int       `json:"poin" gorm:"not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	HasilTestInputParam struct {
		Poin int `json:"poin" binding:"required"`
	}

	HasilTestResponse struct {
		ID        uint   `json:"id"`
		Poin      int    `json:"poin`
		CreatedAt string `json:"created_at"`
	}
)
