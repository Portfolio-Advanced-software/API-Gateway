package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"github.com/gin-gonic/gin"
)

type UpdateHistoryRequestBody struct {
	UserId   string `json:"userid"`
	MovieId  string `json:"movieid"`
	Progress string `json:"progress"`
	Like     bool   `json:"like"`
}

func UpdateHistory(ctx *gin.Context, c pb.HistoryServiceClient) {
	id := ctx.Param("id")

	// Bind the request body to the UpdateHistoryRequestBody struct
	body := UpdateHistoryRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create a gRPC request message with the updated History information
	updateReq := &pb.UpdateHistoryReq{
		History: &pb.History{
			Id:       id,
			Userid:   body.UserId,
			Movieid:  body.MovieId,
			Progress: body.Progress,
			Like:     body.Like,
		},
	}

	// Call the UpdateHistory RPC
	res, err := c.UpdateHistory(context.Background(), updateReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
