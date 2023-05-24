package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"github.com/gin-gonic/gin"
)

func ListHistories(ctx *gin.Context, c pb.HistoryServiceClient) {
	// Call the ListHistories RPC
	res, err := c.ListHistories(context.Background(), &pb.ListHistoriesReq{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
