package movie

import (
	"github.com/gin-gonic/gin"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/movie")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateMovie)
	routes.GET("/:id", svc.ReadMovie)
}

func (svc *ServiceClient) CreateMovie(ctx *gin.Context) {
	routes.CreateMovie(ctx, svc.Client)
}

func (svc *ServiceClient) ReadMovie(ctx *gin.Context) {
	routes.ReadMovie(ctx, svc.Client)
}
