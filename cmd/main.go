package main

import (
	"log"

	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/auth"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/authz"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/config"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/history"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/movie"
	promotheus "github.com/Portfolio-Advanced-software/API-Gateway/pkg/prometheus"
	"github.com/Portfolio-Advanced-software/API-Gateway/pkg/user"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	authzSvc := *authz.RegisterRoutes(r, &c)

	movie.RegisterRoutes(r, &c, &authSvc, &authzSvc)
	user.RegisterRoutes(r, &c, &authSvc)
	history.RegisterRoutes(r, &c, &authSvc)
	promotheus.RegisterRoutes(r, &c)

	r.Run(c.Port)

}
