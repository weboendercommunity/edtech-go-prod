package admin

import (
	adminEntity "edtech.id/internal/admin/entity"
	"edtech.id/pkg/utils"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll(offset int, limit int) []adminEntity.Admin
	FindById(id int) (*adminEntity.Admin, error)
	FindByEmail(email string) (*adminEntity.Admin, error)
	Create(adminEntity adminEntity.Admin) (*adminEntity.Admin, error)
	Update(adminEntity adminEntity.Admin) (*adminEntity.Admin, error)
	Delete(adminEntity adminEntity.Admin) error
}

type AdminRepositoryImpl struct {
	db *gorm.DB
}

// Create implements AdminRepository
func (ar *AdminRepositoryImpl) Create(adminEntity adminEntity.Admin) (*adminEntity.Admin, error) {
	createdAdmin := ar.db.Create(&adminEntity)

	if createdAdmin.Error != nil {
		return nil, createdAdmin.Error
	}

	return &adminEntity, nil
}

// FindAll implements AdminRepository
func (ar *AdminRepositoryImpl) FindAll(offset int, limit int) []adminEntity.Admin {
	var admins []adminEntity.Admin

	ar.db.Scopes(utils.Paginate(offset, limit)).Find(&admins)

	return admins
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
func (ar *AdminRepositoryImpl) FindById(id int) (*adminEntity.Admin, error) {
	var admin adminEntity.Admin

	dataAdmin := ar.db.First(&admin, id)

	if dataAdmin.Error != nil {
		return nil, dataAdmin.Error
	}

	return &admin, nil
}

// Update implements AdminRepository
func (ar *AdminRepositoryImpl) Update(adminEntity adminEntity.Admin) (*adminEntity.Admin, error) {
	updatedAdmin := ar.db.Save(&adminEntity)

	if updatedAdmin.Error != nil {
		return nil, updatedAdmin.Error
	}

	return &adminEntity, nil
}

// Delete implements AdminRepository
func (ar *AdminRepositoryImpl) Delete(adminEntity adminEntity.Admin) error {
	deletedAdmin := ar.db.Delete(&adminEntity)

	if deletedAdmin.Error != nil {
		return deletedAdmin.Error
	}

	return nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{db}
}
