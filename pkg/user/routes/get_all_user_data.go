package routes

import (
	"context"
	"net/http"

	userpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func GetAllUserData(ctx *gin.Context, c userpb.UserServiceClient) {
	// Get the object ID from the request parameters or body
	objectID := ctx.Param("id") // Example: assuming the object ID is part of the URL path

	// Call the DeleteUser RPC
	res, err := c.GetAllUserData(context.Background(), &userpb.GetAllUserDataReq{
		Id: objectID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
