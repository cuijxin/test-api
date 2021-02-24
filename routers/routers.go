package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	healthcheck "github.com/RaMin0/gin-health-check"

	v1 "github.com/cuijxin/mysql-cluster-presslabs/routers/api"
	"github.com/cuijxin/mysql-cluster-presslabs/routers/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(healthcheck.Default())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	check := g.Group("/healthcheck")
	{
		check.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	apiv1 := g.Group("/api/v1")
	{
		apiv1.GET("/mysqlclusters", v1.GetMysqlClusters)
	}

	return g
}
