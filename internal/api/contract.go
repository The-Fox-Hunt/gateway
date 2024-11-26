package api

import "context"

type AuthClient interface {
	DoSignUp(ctx context.Context) error
}
