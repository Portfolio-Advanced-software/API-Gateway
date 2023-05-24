package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func ListUsers(ctx *gin.Context, c pb.UserServiceClient) {
	// Call the ListUsers RPC
	res, err := c.ListUsers(context.Background(), &pb.ListUsersReq{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
