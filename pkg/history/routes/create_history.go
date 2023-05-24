package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"github.com/gin-gonic/gin"
)

type CreateHistoryRequestBody struct {
	UserId   string `json:"userid"`
	MovieId  string `json:"movieid"`
	Progress string `json:"progress"`
	Like     bool   `json:"like"`
}

func CreateHistory(ctx *gin.Context, c pb.HistoryServiceClient) {
	body := CreateHistoryRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateHistory(context.Background(), &pb.CreateHistoryReq{
		History: &pb.History{
			Userid:   body.UserId,
			Movieid:  body.MovieId,
			Progress: body.Progress,
			Like:     body.Like,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
