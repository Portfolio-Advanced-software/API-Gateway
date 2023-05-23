package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"github.com/gin-gonic/gin"
)

type UpdateHistoryRequestBody struct {
	UserId   int64  `json:"userid"`
	MovieId  int64  `json:"movieid"`
	Progress string `json:"progress"`
	Like     bool   `json:"like"`
}

func UpdateHistory(ctx *gin.Context, c pb.HistoryServiceClient) {
	userid := ctx.Param("userid")
	movieid := ctx.Param("movieid")

	// Bind the request body to the UpdateMovieRequestBody struct
	body := UpdateHistoryRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create a gRPC request message with the updated user information
	updateReq := &pb.UpdateUserReq{
		User: &pb.User{
			Userid:   userid,
			Movieid:  movieid,
			Progress: body.Progress,
			Like:     body.Like,
		},
	}

	// Call the UpdateUser RPC
	res, err := c.UpdateHistory(context.Background(), updateReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
