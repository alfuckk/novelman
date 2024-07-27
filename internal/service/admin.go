package service

import (
	"context"
	v1 "novelman/api/v1"
	"novelman/internal/model"
	"novelman/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	GetAdmin(ctx context.Context, id int64) (*model.Admin, error)
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
}

func NewAdminService(
	service *Service,
	adminRepository repository.AdminRepository,
) AdminService {
	return &adminService{
		Service:         service,
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

func (s *adminService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	admin, err := s.adminRepository.GetByEmail(ctx, req.Email)
	if err != nil || admin == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(admin.AdminID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}
