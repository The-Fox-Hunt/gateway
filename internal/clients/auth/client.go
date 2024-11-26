package auth

import (
	"context"
	"fmt"
	"github.com/The-Fox-Hunt/auth/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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

func (c *Client) DoSignUp(ctx context.Context) error {
	resp, err := c.client.Login(ctx, &auth.LoginIn{
		Username: "Hello ",
		Password: "World",
	})
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	log.Printf("login succeeded: %s", resp.Token)
	return nil
}
