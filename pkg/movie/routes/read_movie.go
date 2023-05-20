package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
)

func ReadMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.ReadMovie(context.Background(), &pb.ReadMovieReq{
		Id: strconv.FormatInt(id, 10),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
