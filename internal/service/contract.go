package service

import (
	"context"
	"github.com/The-Fox-Hunt/gateway/internal/model"
)

type AuthClient interface {
	DoSignUp(ctx context.Context, data model.SignupData) (model.SignupSuccess, error)
}
