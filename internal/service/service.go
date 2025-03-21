package service

import (
	"context"
	"fmt"

	"github.com/The-Fox-Hunt/gateway/internal/model"
	"google.golang.org/grpc/metadata"
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

func (s *Service) SignIn(ctx context.Context, data model.SignInData) (model.SignInSuccess, error) {
	res, err := s.authC.DoSignIn(ctx, data)
	if err != nil {
		return model.SignInSuccess{}, fmt.Errorf("failed to make requst for signip: %w", err)
	}
	return res, nil
}

func (s *Service) ChangePassword(ctx context.Context, data model.ChangePasswordData) (model.ChangePasswordSuccess, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("username", ctx.Value(model.Username).(string)))
	// Отправляем gRPC-запрос в `auth`
	res, err := s.authC.DoChangePassword(ctx, data)
	if err != nil {
		return model.ChangePasswordSuccess{}, fmt.Errorf("failed to make request for changepassword: %w", err)
	}

	return model.ChangePasswordSuccess{Success: res.Success}, nil
}
