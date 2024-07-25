package service

import (
	"context"
	v1 "novelman/api/v1"
	"novelman/internal/model"
	"novelman/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
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

func (s *adminService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check adminname
	admin, err := s.adminRepository.GetByEmail(ctx, req.Email)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if err == nil && admin != nil {
		return v1.ErrEmailAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate admin ID
	adminID, err := s.sid.GenString()
	if err != nil {
		return err
	}
	admin = &model.Admin{
		AdminID:  adminID,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a admin
		if err = s.adminRepository.Create(ctx, admin); err != nil {
			return err
		}
		// TODO: other repo
		return nil
	})
	return err
}

func (s *adminService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {

	return "", nil
}
