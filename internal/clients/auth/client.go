package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/The-Fox-Hunt/auth/pkg/auth"
	"github.com/The-Fox-Hunt/gateway/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client auth.AuthServiceClient
}

func NewClient() *Client {
	conn, err := grpc.NewClient("auth:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := auth.NewAuthServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) DoSignUp(ctx context.Context, data model.SignupData) (model.SignupSuccess, error) {
	resp, err := c.client.Signup(ctx, &auth.SignupIn{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		return model.SignupSuccess{}, fmt.Errorf("failed to login: %w", err)
	}

	log.Printf("login succeeded: %t", resp.Success)
	return model.SignupSuccess{
		Success: resp.Success,
	}, nil
}

func (c *Client) DoSignIn(ctx context.Context, data model.SignInData) (model.SignInSuccess, error) {
	resp, err := c.client.Login(ctx, &auth.LoginIn{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		return model.SignInSuccess{}, fmt.Errorf("failed to login: %w", err)
	}

	return model.SignInSuccess{
		Token: resp.Token,
	}, nil
}

func (c *Client) DoChangePassword(ctx context.Context, data model.ChangePasswordData) (model.ChangePasswordSuccess, error) {
	username, ok := ctx.Value(model.Username).(string)
	if !ok || username == "" {
		return model.ChangePasswordSuccess{}, fmt.Errorf("username not found in context")
	}

	// Отправляем запрос в сервис `auth`
	resp, err := c.client.ChangePassword(ctx, &auth.ChangePasswordIn{
		OldPassword: data.OldPassword,
		NewPassword: data.NewPassword,
	})

	if err != nil {
		return model.ChangePasswordSuccess{}, fmt.Errorf("failed to change password: %w", err)
	}

	return model.ChangePasswordSuccess{Success: resp.Success}, nil
}
