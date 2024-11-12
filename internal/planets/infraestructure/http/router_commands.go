package http

import (
	"github.com/gin-gonic/gin"
	"github.com/id4n4/go_cqrs/internal/planets/app/commands"
)

func SetupCommandsRouter(createHandler *commands.CreatePlanetCommandHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/planets", func(c *gin.Context) {
		var cmd commands.CreatePlanetCommand
		if err := c.ShouldBindJSON(&cmd); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := createHandler.HandleCreate(c.Request.Context(), cmd); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{
			"message": "Planet created",
		})
	})

	return r
}
