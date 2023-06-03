package routes

import (
	"context"
	"net/http"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz/pb"
	"github.com/gin-gonic/gin"
)

type VerifyRoleRequestBody struct {
	JwtToken string `json:"jwt_token"`
}

func VerifyRole(ctx *gin.Context, c pb.AuthzServiceClient) {
	b := VerifyRoleRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.VerifyRole(context.Background(), &pb.VerifyRoleRequest{
		JwtToken: b.JwtToken,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
