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
	conn, err := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	log.Printf("login succeeded: %s", resp.Success)
	return model.SignupSuccess{
		Success: resp.Success,
	}, nil
}

func (c *Client) DoSignIn(ctx context.Context, data model.SignInData) (model.SignInSucess, error) {
	resp, err := c.client.Login(ctx, &auth.LoginIn{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		return model.SignInSucess{}, fmt.Errorf("failed to login: %w", err)
	}

	return model.SignInSucess{
		Token: resp.Token,
	}, nil
}
