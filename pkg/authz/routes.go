package authz

import (
	"github.com/gin-gonic/gin"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz/routes"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/authz")
	routes.POST("/verify", svc.VerifyRole)

	return svc
}

func (svc *ServiceClient) VerifyRole(ctx *gin.Context) {
	routes.VerifyRole(ctx, svc.Client)
}
