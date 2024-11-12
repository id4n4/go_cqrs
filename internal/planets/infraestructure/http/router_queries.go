package http

import (
	"github.com/gin-gonic/gin"
	"github.com/id4n4/go_cqrs/internal/planets/app/queries"
)

type GetPlanetParams struct {
	ID string `uri:"id" binding:"required"`
}

func SetupQueriesRouter(getHandler *queries.GetPlanetQueryHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/planets", func(c *gin.Context) {
		planets, err := getHandler.HandleGetAll(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, planets)
	})

	r.GET("/planets/:id", func(c *gin.Context) {

		var params GetPlanetParams

		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		planet, err := getHandler.HandleGet(c.Request.Context(), queries.GetPlanetQuery{ID: params.ID})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, planet)
	})

	return r
}
