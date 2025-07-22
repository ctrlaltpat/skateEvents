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
	{
		v1.POST("/events", app.handlers.Event.CreateEvent)
		v1.GET("/events", app.handlers.Event.GetAllEvents)
		v1.GET("/events/:id", app.handlers.Event.GetEvent)
		v1.PUT("/events/:id", app.handlers.Event.UpdateEvent)
		v1.DELETE("/events/:id", app.handlers.Event.DeleteEvent)
		v1.PATCH("/events/:id/status", app.handlers.Event.UpdateEventStatus)

		v1.POST("/events/:id/rsvp/:userId", app.handlers.Event.AddAttendeeToEvent)
		v1.GET("/events/:id/attendees", app.handlers.Event.GetAttendeesByEventId)
	}

	v1.POST("/auth/register", app.handlers.User.RegisterUser)
	v1.GET("/users/:id", app.handlers.User.GetUser)

	return g
}
