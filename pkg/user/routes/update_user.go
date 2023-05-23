package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequestBody struct {
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	DateOfBirth      string `json:"dateofbirth"`
	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	CreditCardNumber int32  `json:"creditcardnumber"`
	ExpirationDate   string `json:"expirationdate"`
	CVC              int32  `json:"cvc"`
}

func UpdateUser(ctx *gin.Context, c pb.UserServiceClient) {
	id := ctx.Param("id")

	// Bind the request body to the UpdateMovieRequestBody struct
	body := UpdateUserRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create a gRPC request message with the updated user information
	updateReq := &pb.UpdateUserReq{
		User: &pb.User{
			Id:               id,
			Email:            body.Email,
			Phone:            body.Phone,
			Dateofbirth:      body.DateOfBirth,
			Firstname:        body.FirstName,
			Lastname:         body.LastName,
			Creditcardnumber: body.CreditCardNumber,
			Expirationdate:   body.ExpirationDate,
			Cvc:              body.CVC,
		},
	}

	// Call the UpdateUser RPC
	res, err := c.UpdateUser(context.Background(), updateReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
