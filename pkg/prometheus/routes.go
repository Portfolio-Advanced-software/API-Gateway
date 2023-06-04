package promotheus

import (
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
