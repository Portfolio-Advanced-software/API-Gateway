package history

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient, authzSvc *authz.ServiceClient, svc *ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	az := authz.InitAuthzMiddleware(authzSvc)

	// svc := &ServiceClient{
	// 	Client: InitServiceClient(c),
	// }

	routes := r.Group("/history")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateHistory)
	routes.PUT("/:id", svc.UpdateHistory)
	routes.DELETE("/:id", svc.DeleteHistory)

	adminRoutes := r.Group("/history")
	adminRoutes.Use(a.AuthRequired)
	adminRoutes.Use(az.RoleRequired("admin"))
	adminRoutes.GET("/", svc.ListHistories)
	adminRoutes.GET("/:id", svc.ReadHistory)
}

func (svc *ServiceClient) CreateHistory(ctx *gin.Context) {
	routes.CreateHistory(ctx, svc.Client)
}

func (svc *ServiceClient) ReadHistory(ctx *gin.Context) {
	fmt.Println("Handling ReadHistory request")
	routes.ReadHistory(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateHistory(ctx *gin.Context) {
	routes.UpdateHistory(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteHistory(ctx *gin.Context) {
	routes.DeleteHistory(ctx, svc.Client)
}

func (svc *ServiceClient) ListHistories(ctx *gin.Context) {
	routes.ListHistories(ctx, svc.Client)
}
