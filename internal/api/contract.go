package api

import (
	"context"

	"github.com/The-Fox-Hunt/gateway/internal/model"
)

type AuthService interface {
	SignUp(ctx context.Context, data model.SignupData) (model.SignupSuccess, error)
	SignIn(ctx context.Context, data model.SignInData) (model.SignInSucess, error)
}
