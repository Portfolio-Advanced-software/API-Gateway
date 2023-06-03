package movie

import (
	"github.com/gin-gonic/gin"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient, authzSvc *authz.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	az := authz.InitAuthzMiddleware(authzSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/movie")
	routes.Use(a.AuthRequired)
	routes.Use(az.AuthzRequired("user"))

	routes.POST("/", svc.CreateMovie)
	routes.GET("/:id", svc.ReadMovie)
	routes.PUT("/:id", svc.UpdateMovie)
	routes.DELETE("/:id", svc.DeleteMovie)
	routes.GET("/", svc.ListMovies)
}

func (svc *ServiceClient) CreateMovie(ctx *gin.Context) {
	routes.CreateMovie(ctx, svc.Client)
}

func (svc *ServiceClient) ReadMovie(ctx *gin.Context) {
	routes.ReadMovie(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateMovie(ctx *gin.Context) {
	routes.UpdateMovie(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteMovie(ctx *gin.Context) {
	routes.DeleteMovie(ctx, svc.Client)
}

func (svc *ServiceClient) ListMovies(ctx *gin.Context) {
	routes.ListMovies(ctx, svc.Client)
}
