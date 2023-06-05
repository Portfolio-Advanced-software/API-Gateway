package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	id := ctx.Param("id")

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(objectID)
	res, err := c.ReadMovie(context.Background(), &pb.ReadMovieReq{
		Id: objectID.Hex(),
	})

	if err != nil {
		fmt.Println("here")
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
