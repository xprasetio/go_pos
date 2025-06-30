package user_domain

import (
	"context"
	"pos-go/pkg/jwt"
	"pos-go/pkg/response"
)

type UserService interface {
	Register(ctx context.Context, name, email, password string, roleID int64) error
	Login(ctx *LoginRequest) (*User, string, error)
	FindByID(id uint) (*User, error)
}

type userService struct {
	repo UserRepository
	jwt  jwt.JWTService
}

func NewUserService(repo UserRepository, jwtSvc jwt.JWTService) UserService {
	return &userService{
		repo: repo,
		jwt:  jwtSvc,
	}
}

func (s *userService) Register(ctx context.Context, name, email, password string, roleID int64) error {
	hash, err := response.HashPassword(password)
	if err != nil {
		return err
	}
	user := &User{
		Name:     name,
		Email:    email,
		Password: hash,
		RoleID:   roleID,
	}
	return s.repo.Create(ctx, user)
}

func (s *userService) Login(req *LoginRequest) (*User, string, error) {
	user, err := s.repo.FindByEmail(context.Background(), req.Email)
	if err != nil {
		return nil, "", err
	}
	if !response.CheckPasswordHash(req.Password, user.Password) {
		return nil, "", err
	}
	token, err := s.jwt.GenerateToken(uint(user.ID), user.Email)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}

func (s *userService) FindByID(id uint) (*User, error) {
	return s.repo.FindByID(context.Background(), int64(id))
}
