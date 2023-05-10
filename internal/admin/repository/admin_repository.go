package admin

import (
	adminEntity "edtech.id/internal/admin/entity"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll(offset int, limit int) []*adminEntity.Admin
	FindById(id int) (*adminEntity.Admin, error)
	FindByEmail(email string) (*adminEntity.Admin, error)
	Create(adminEntity adminEntity.Admin) (*adminEntity.Admin, error)
	Update(adminEntity adminEntity.Admin) (*adminEntity.Admin, error)
}

type AdminRepositoryImpl struct {
	db *gorm.DB
}

// Create implements AdminRepository
func (*AdminRepositoryImpl) Create(adminEntity adminEntity.Admin) (*adminEntity.Admin, error) {
	panic("unimplemented")
}

// FindAll implements AdminRepository
func (*AdminRepositoryImpl) FindAll(offset int, limit int) []*adminEntity.Admin {
	panic("unimplemented")
}

// FindByEmail implements AdminRepository
func (ar *AdminRepositoryImpl) FindByEmail(email string) (*adminEntity.Admin, error) {
	var admin adminEntity.Admin

	admins := ar.db.Where("email = ?", email).First(&admin)

	if admins.Error != nil {
		return nil, admins.Error
	}

	return &admin, nil
}

// FindById implements AdminRepository
func (*AdminRepositoryImpl) FindById(id int) (*adminEntity.Admin, error) {
	panic("unimplemented")
}

// Update implements AdminRepository
func (*AdminRepositoryImpl) Update(adminEntity adminEntity.Admin) (*adminEntity.Admin, error) {
	panic("unimplemented")
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{db}
}
