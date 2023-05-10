package admin

import (
	"database/sql"

	"gorm.io/gorm"
)

type Admin struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"-"`
	CreatedBy   *Admin         `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64         `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   *Admin         `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
