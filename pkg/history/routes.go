package history

import (
	"github.com/gin-gonic/gin"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	//a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/history")
	//routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateHistory)
	routes.GET("/:id", svc.ReadHistory)
	routes.PUT("/:id", svc.UpdateHistory)
	routes.DELETE("/:id", svc.DeleteHistory)
	routes.GET("/", svc.ListHistories)
}

func (svc *ServiceClient) CreateHistory(ctx *gin.Context) {
	routes.CreateHistory(ctx, svc.Client)
}

func (svc *ServiceClient) ReadHistory(ctx *gin.Context) {
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
