package admin

import "database/sql"

type Admin struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	CreatedBy   *Admin       `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64       `json:"created_by"`
	UpdatedBy   *Admin       `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID *int64       `json:"updated_by"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}
