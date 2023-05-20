package movie

import (
	"fmt"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.MovieServiceClient
}

func InitServiceClient(c *config.Config) pb.MovieServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.MovieSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewMovieServiceClient(cc)
}
