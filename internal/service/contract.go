package service

//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go

import (
	"context"

	"github.com/The-Fox-Hunt/gateway/internal/model"
)

type AuthClient interface {
	DoSignUp(ctx context.Context, data model.SignupData) (model.SignupSuccess, error)
	DoSignIn(ctx context.Context, data model.SignInData) (model.SignInSuccess, error)
	DoChangePassword(ctx context.Context, data model.ChangePasswordData) (model.ChangePasswordSuccess, error)
}
