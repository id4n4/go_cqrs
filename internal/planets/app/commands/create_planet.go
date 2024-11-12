package commands

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/id4n4/go_cqrs/internal/planets/domain"
)

type CreatePlanetCommand struct {
	Name         string             `json:"name" binding:"required"`
	Resource     string             `json:"resource" binding:"required"`
	PlanetType   domain.PlanetTypes `json:"planet_type" binding:"required"`
	Distance     int                `json:"distance" binding:"required"`
	OrbitalOrder int                `json:"orbital_order" binding:"required"`
}

type CreatePlanetCommandHandler struct {
	repo domain.Repository
}

func NewCreatePlanetCommandHandler(repo domain.Repository) *CreatePlanetCommandHandler {
	return &CreatePlanetCommandHandler{repo}
}

func (h *CreatePlanetCommandHandler) HandleCreate(ctx context.Context, cmd CreatePlanetCommand) error {
	planet := &domain.Planet{
		ID:           uuid.New().String(),
		Name:         cmd.Name,
		Resource:     cmd.Resource,
		PlanetType:   cmd.PlanetType,
		Distance:     cmd.Distance,
		OrbitalOrder: cmd.OrbitalOrder,
		CreatedAt:    time.Now(),
	}

	return h.repo.CreatePlanet(ctx, planet)
}
