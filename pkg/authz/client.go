package authz

import (
	"fmt"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz/pb"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthzServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthzServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthzSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthzServiceClient(cc)
}
