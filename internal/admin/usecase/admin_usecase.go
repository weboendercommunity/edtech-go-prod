package admin

import (
	adminDto "edtech.id/internal/admin/dto"
	adminEntity "edtech.id/internal/admin/entity"
	adminRepository "edtech.id/internal/admin/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	FindAll(offset int, limit int) []adminEntity.Admin
	FindById(id int) (*adminEntity.Admin, error)
	FindByEmail(email string) (*adminEntity.Admin, error)
	Create(adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error)
	Update(id int, adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error)
	Delete(id int) error
}

type AdminUseCaseImpl struct {
	adminRepository adminRepository.AdminRepository
}

// Create implements AdminUseCase
func (au *AdminUseCaseImpl) Create(adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*adminDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dataAdmin := adminEntity.Admin{
		Name:        adminDto.Name,
		Email:       adminDto.Email,
		Password:    string(hashedPassword),
		CreatedByID: &adminDto.CreatedBy,
	}

	admin, err := au.adminRepository.Create(dataAdmin)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

// Delete implements AdminUseCase
func (au *AdminUseCaseImpl) Delete(id int) error {
	admin, err := au.adminRepository.FindById(id)

	if err != nil {
		return err
	}

	deletedAdmin := au.adminRepository.Delete(*admin)

	if deletedAdmin != nil {
		return deletedAdmin
	}

	return nil
}

// FindAll implements AdminUseCase
func (au *AdminUseCaseImpl) FindAll(offset int, limit int) []adminEntity.Admin {
	return au.adminRepository.FindAll(offset, limit)
}

// FindByEmail implements AdminUseCase
func (au *AdminUseCaseImpl) FindByEmail(email string) (*adminEntity.Admin, error) {
	return au.adminRepository.FindByEmail(email)
}

// FindById implements AdminUseCase
func (au *AdminUseCaseImpl) FindById(id int) (*adminEntity.Admin, error) {
	return au.adminRepository.FindById(id)
}

// Update implements AdminUseCase
func (au *AdminUseCaseImpl) Update(id int, adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error) {
	admin, err := au.adminRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	admin.Name = adminDto.Name

	//  validate email is different with current email
	if admin.Email != adminDto.Email {
		admin.Email = adminDto.Email
	}

	if adminDto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*adminDto.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}

		admin.Password = string(hashedPassword)
	}

	admin.UpdatedByID = &adminDto.UpdatedBy

	// update admin

	updatedAdmin, err := au.adminRepository.Update(*admin)

	if err != nil {
		return nil, err
	}

	return updatedAdmin, nil

}

func NewAdminUseCase(adminRepository adminRepository.AdminRepository) AdminUseCase {
	return &AdminUseCaseImpl{adminRepository}
}
