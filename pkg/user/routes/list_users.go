package routes

import (
	"context"
	"net/http"

	userpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func ListUsers(ctx *gin.Context, c userpb.UserServiceClient) {
	// Call the ListUsers RPC
	res, err := c.ListUsers(context.Background(), &userpb.ListUsersReq{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
