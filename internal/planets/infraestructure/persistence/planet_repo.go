package persistence

import (
	"context"
	"database/sql"

	"github.com/id4n4/go_cqrs/internal/planets/domain"
	_ "github.com/lib/pq"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{db}
}

func (r *UserRepositoryPostgres) Close() {
	r.db.Close()
}

func (r *UserRepositoryPostgres) GetPlanets(ctx context.Context) ([]*domain.Planet, error) {
	query := `SELECT id, name, resource, planet_type, distance, "orbital_order", created_at FROM planets`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planets []*domain.Planet
	for rows.Next() {
		var planet domain.Planet
		if err := rows.Scan(&planet.ID, &planet.Name, &planet.Resource, &planet.PlanetType, &planet.Distance, &planet.OrbitalOrder, &planet.CreatedAt); err != nil {
			return nil, err
		}
		planets = append(planets, &planet)
	}

	return planets, nil
}

func (r *UserRepositoryPostgres) GetPlanetByID(ctx context.Context, id string) (*domain.Planet, error) {
	query := `SELECT id, name, resource, planet_type, distance, "orbital_order", created_at FROM planets WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var planet domain.Planet
	if err := row.Scan(&planet.ID, &planet.Name, &planet.Resource, &planet.PlanetType, &planet.Distance, &planet.OrbitalOrder, &planet.CreatedAt); err != nil {
		return nil, err
	}

	return &planet, nil
}

func (r *UserRepositoryPostgres) CreatePlanet(ctx context.Context, planet *domain.Planet) error {
	query := `INSERT INTO planets (id, name, resource, planet_type, distance, "orbital_order", created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, planet.ID, planet.Name, planet.Resource, planet.PlanetType, planet.Distance, planet.OrbitalOrder, planet.CreatedAt)
	return err
}

func (r *UserRepositoryPostgres) UpdatePlanet(ctx context.Context, planet *domain.Planet) error {
	query := `UPDATE planets SET name = $1, resource = $2, planet_type = $3, distance = $4, "orbital_order" = $5, created_at = $6 WHERE id = $7`
	_, err := r.db.ExecContext(ctx, query, planet.Name, planet.Resource, planet.PlanetType, planet.Distance, planet.OrbitalOrder, planet.CreatedAt, planet.ID)
	return err
}

func (r *UserRepositoryPostgres) DeletePlanet(ctx context.Context, id string) error {
	query := `DELETE FROM planets WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
