package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"github.com/gin-gonic/gin"
)

type CreateHistoryRequestBody struct {
	UserId   int64  `json:"userid"`
	MovieId  int64  `json:"movieid"`
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
			UserId:   body.UserId,
			MovieId:  body.MovieId,
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
