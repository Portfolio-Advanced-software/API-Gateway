package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
)

type UpdateMovieRequestBody struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Director    string  `json:"director"`
	Genre       string  `json:"genre"`
	Rating      float32 `json:"rating"`
	Runtime     int32   `json:"runtime"`
	Poster      string  `json:"poster"`
}

func UpdateMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	id := ctx.Param("id")

	// Bind the request body to the UpdateMovieRequestBody struct
	body := UpdateMovieRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create a gRPC request message with the updated movie information
	updateReq := &pb.UpdateMovieReq{
		Movie: &pb.Movie{
			Id:          id,
			Title:       body.Title,
			Description: body.Description,
			ReleaseDate: body.ReleaseDate,
			Director:    body.Director,
			Genre:       body.Genre,
			Rating:      body.Rating,
			Runtime:     body.Runtime,
			Poster:      body.Poster,
		},
	}

	// Call the UpdateMovie RPC
	res, err := c.UpdateMovie(context.Background(), updateReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
