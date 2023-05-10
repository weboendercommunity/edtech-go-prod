package admin

import (
	adminDto "edtech.id/internal/admin/dto"
	adminEntity "edtech.id/internal/admin/entity"
	adminRepository "edtech.id/internal/admin/repository"
)

type AdminUseCase interface {
	FindAll(offset int, limit int) []adminEntity.Admin
	FindById(id int) (*adminEntity.Admin, error)
	FindByEmail(email string) (*adminEntity.Admin, error)
	Create(adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error)
	Update(adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error)
	Delete(id int) error
}

type AdminUseCaseImpl struct {
	adminRepository adminRepository.AdminRepository
}

// Create implements AdminUseCase
func (*AdminUseCaseImpl) Create(adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error) {
	panic("unimplemented")
}

// Delete implements AdminUseCase
func (*AdminUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements AdminUseCase
func (*AdminUseCaseImpl) FindAll(offset int, limit int) []adminEntity.Admin {
	panic("unimplemented")
}

// FindByEmail implements AdminUseCase
func (au *AdminUseCaseImpl) FindByEmail(email string) (*adminEntity.Admin, error) {
	return au.adminRepository.FindByEmail(email)
}

// FindById implements AdminUseCase
func (*AdminUseCaseImpl) FindById(id int) (*adminEntity.Admin, error) {
	panic("unimplemented")
}

// Update implements AdminUseCase
func (*AdminUseCaseImpl) Update(adminDto adminDto.AdminRequestBody) (*adminEntity.Admin, error) {
	panic("unimplemented")
}

func NewAdminUseCase(adminRepository adminRepository.AdminRepository) AdminUseCase {
	return &AdminUseCaseImpl{adminRepository}
}
