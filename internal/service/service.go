package service

import (
	"context"
	"fmt"
	"github.com/The-Fox-Hunt/gateway/internal/model"
)

type Service struct {
	authC AuthClient
}

func NewService(aC AuthClient) *Service {
	return &Service{
		authC: aC,
	}
}

func (s *Service) SignUp(ctx context.Context, data model.SignupData) (model.SignupSuccess, error) {
	if data.Password != data.ConfirmPassword {
		return model.SignupSuccess{}, fmt.Errorf("password does not match")
	}
	res, err := s.authC.DoSignUp(ctx, data)
	if err != nil {
		return model.SignupSuccess{}, fmt.Errorf("failed to make requst for signup: %w", err)
	}
	return res, nil
}
