package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"github.com/gin-gonic/gin"
)

func DeleteHistory(ctx *gin.Context, c pb.HistoryServiceClient) {
	// Get the object ID from the request parameters or body
	objectID := ctx.Param("id") // Example: assuming the object ID is part of the URL path

	// Call the DeleteHistory RPC
	res, err := c.DeleteHistory(context.Background(), &pb.DeleteHistoryReq{
		Id: objectID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
