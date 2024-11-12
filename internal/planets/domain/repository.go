package domain

import "context"

type Repository interface {
	Close()
	GetPlanets(ctx context.Context) ([]*Planet, error)
	GetPlanetByID(ctx context.Context, id string) (*Planet, error)
	CreatePlanet(ctx context.Context, planet *Planet) error
	UpdatePlanet(ctx context.Context, planet *Planet) error
	DeletePlanet(ctx context.Context, id string) error
}

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func Close() {
	repository.Close()
}

func GetPlanets(ctx context.Context) ([]*Planet, error) {
	return repository.GetPlanets(ctx)
}

func GetPlanetByID(ctx context.Context, id string) (*Planet, error) {
	return repository.GetPlanetByID(ctx, id)
}

func CreatePlanet(ctx context.Context, planet *Planet) error {
	return repository.CreatePlanet(ctx, planet)
}

func UpdatePlanet(ctx context.Context, planet *Planet) error {
	return repository.UpdatePlanet(ctx, planet)
}

func DeletePlanet(ctx context.Context, id string) error {
	return repository.DeletePlanet(ctx, id)
}
