package main

import (
	"github.com/id4n4/go_cqrs/internal/planets/app/commands"
	"github.com/id4n4/go_cqrs/internal/planets/infraestructure/config"
	"github.com/id4n4/go_cqrs/internal/planets/infraestructure/http"
	"github.com/id4n4/go_cqrs/internal/planets/infraestructure/persistence"
)

func main() {
	db := config.ConnectDB()

	defer db.Close()

	planetsRepo := persistence.NewUserRepositoryPostgres(db)
	createPlanetsHandler := commands.NewCreatePlanetCommandHandler(planetsRepo)

	router := http.SetupCommandsRouter(createPlanetsHandler)
	router.Run(":8080")
}
