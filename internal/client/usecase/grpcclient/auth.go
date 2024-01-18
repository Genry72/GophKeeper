package grpcclient

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/proto"
	"google.golang.org/grpc/status"
)

func (c *Client) Auth(ctx context.Context, username, password string) (string, error) {
	authMsg := &proto.AuthUserMsg{
		Username: username,
		Password: password,
	}

	token, err := c.usersClient.Auth(ctx, authMsg)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return "", fmt.Errorf(e.Message())
		}

	}

	return token.Token, nil
}
