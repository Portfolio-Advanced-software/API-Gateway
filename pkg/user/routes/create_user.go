package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type CreateUserRequestBody struct {
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	DateOfBirth      string `json:"dateofbirth"`
	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	CreditCardNumber int32  `json:"creditcardnumber"`
	ExpirationDate   string `json:"expirationdate"`
	CVC              int32  `json:"cvc"`
}

func CreateUser(ctx *gin.Context, c pb.UserServiceClient) {
	body := CreateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateUser(context.Background(), &pb.CreateUserReq{
		User: &pb.User{
			Email:            body.Email,
			Phone:            body.Phone,
			Dateofbirth:      body.DateOfBirth,
			Firstname:        body.FirstName,
			Lastname:         body.LastName,
			Creditcardnumber: body.CreditCardNumber,
			Expirationdate:   body.ExpirationDate,
			Cvc:              body.CVC,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
