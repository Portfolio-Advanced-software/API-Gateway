package routes

import (
	"context"
	"net/http"

	userpb "github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type CreateUserRequestBody struct {
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	DateOfBirth      string `json:"date_of_birth"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	CreditCardNumber int32  `json:"credit_card_number"`
	ExpirationDate   string `json:"expiration_date"`
	CVC              int32  `json:"cvc"`
}

func CreateUser(ctx *gin.Context, c userpb.UserServiceClient) {
	body := CreateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateUser(context.Background(), &userpb.CreateUserReq{
		User: &userpb.User{
			Email:            body.Email,
			Phone:            body.Phone,
			DateOfBirth:      body.DateOfBirth,
			FirstName:        body.FirstName,
			LastName:         body.LastName,
			CreditCardNumber: body.CreditCardNumber,
			ExpirationDate:   body.ExpirationDate,
			Cvc:              body.CVC,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
