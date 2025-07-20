package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	g.Use(cors.New(config))

	v1 := g.Group("/api/v1")

	v1.POST("/auth/register", app.handlers.User.RegisterUser)

	return g
}
