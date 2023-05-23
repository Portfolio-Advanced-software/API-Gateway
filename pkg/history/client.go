package history

import (
	"fmt"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

func InitServiceClient(c *config.Config) pb.HistoryServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.HistorySvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewUserServiceClient(cc)
}
