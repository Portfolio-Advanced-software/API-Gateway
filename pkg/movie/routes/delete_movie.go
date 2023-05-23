package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
)

func DeleteMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	// Get the object ID from the request parameters or body
	objectID := ctx.Param("id") // Example: assuming the object ID is part of the URL path

	// Call the DeleteMovie RPC
	res, err := c.DeleteMovie(context.Background(), &pb.DeleteMovieReq{
		Id: objectID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
