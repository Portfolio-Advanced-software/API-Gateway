package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context, c pb.MovieServiceClient) {
	// Get the object ID from the request parameters or body
	objectID := ctx.Param("id") // Example: assuming the object ID is part of the URL path

	// Call the DeleteUser RPC
	res, err := c.DeleteUser(context.Background(), &pb.DeleteUserReq{
		Id: objectID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
