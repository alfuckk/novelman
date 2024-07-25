package service

import (
    "context"
	"novelman/internal/model"
	"novelman/internal/repository"
)

type AdminService interface {
	GetAdmin(ctx context.Context, id int64) (*model.Admin, error)
}
func NewAdminService(
    service *Service,
    adminRepository repository.AdminRepository,
) AdminService {
	return &adminService{
		Service:        service,
		adminRepository: adminRepository,
	}
}

type adminService struct {
	*Service
	adminRepository repository.AdminRepository
}

func (s *adminService) GetAdmin(ctx context.Context, id int64) (*model.Admin, error) {
	return s.adminRepository.GetAdmin(ctx, id)
}
