package user

import (
	"fmt"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	userpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client userpb.UserServiceClient
}

func InitServiceClient(c *config.Config) userpb.UserServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.UserSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return userpb.NewUserServiceClient(cc)
}
