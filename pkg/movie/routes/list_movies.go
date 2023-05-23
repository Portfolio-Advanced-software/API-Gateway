package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
)

func ListMovies(ctx *gin.Context, c pb.MovieServiceClient) {
	// Call the ListMovies RPC
	res, err := c.ListMovies(context.Background(), &pb.ListMoviesReq{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
