package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/pb"
	"github.com/gin-gonic/gin"
)

type CreateMovieRequestBody struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Director    string  `json:"director"`
	Genre       string  `json:"genre"`
	Rating      float32 `json:"rating"`
	Runtime     int32   `json:"runtime"`
	Poster      string  `json:"poster"`
}

func CreateMovie(ctx *gin.Context, c pb.MovieServiceClient) {
	body := CreateMovieRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateMovie(context.Background(), &pb.CreateMovieReq{
		Movie: &pb.Movie{
			Title:       body.Title,
			Description: body.Description,
			ReleaseDate: body.ReleaseDate,
			Director:    body.Director,
			Genre:       body.Genre,
			Rating:      body.Rating,
			Runtime:     body.Runtime,
			Poster:      body.Poster,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
