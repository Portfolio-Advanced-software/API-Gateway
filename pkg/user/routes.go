package user

import (
	"github.com/gin-gonic/gin"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/user/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/user")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateUser)
	routes.GET("/:id", svc.ReadUser)
	routes.PUT("/:id", svc.UpdateUser)
	routes.DELETE("/:id", svc.DeleteUser)
	routes.GET("/", svc.ListUsers)
}

func (svc *ServiceClient) CreateUser(ctx *gin.Context) {
	routes.CreateUser(ctx, svc.Client)
}

func (svc *ServiceClient) ReadUser(ctx *gin.Context) {
	routes.ReadUser(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateUser(ctx *gin.Context) {
	routes.UpdateUser(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) {
	routes.DeleteUser(ctx, svc.Client)
}

func (svc *ServiceClient) ListUsers(ctx *gin.Context) {
	routes.ListUsers(ctx, svc.Client)
}
