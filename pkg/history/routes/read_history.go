package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/pb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadHistory(ctx *gin.Context, c pb.HistoryServiceClient) {
	id := ctx.Param("id")

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ReadHistory(context.Background(), &pb.ReadMovieReq{
		Id: objectID.Hex(),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
