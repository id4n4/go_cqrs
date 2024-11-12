package queries

import (
	"context"

	"github.com/id4n4/go_cqrs/internal/planets/domain"
)

type GetPlanetQuery struct {
	ID string `json:"id"`
}

type GetPlanetQueryHandler struct {
	repo domain.Repository
}

func NewGetPlanetQueryHandler(repo domain.Repository) *GetPlanetQueryHandler {
	return &GetPlanetQueryHandler{repo}
}

func (h *GetPlanetQueryHandler) HandleGet(ctx context.Context, query GetPlanetQuery) (*domain.Planet, error) {
	return h.repo.GetPlanetByID(ctx, query.ID)
}

func (h *GetPlanetQueryHandler) HandleGetAll(ctx context.Context) ([]*domain.Planet, error) {
	return h.repo.GetPlanets(ctx)
}
