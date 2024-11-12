package domain

import "time"

type (
	PlanetTypes string
	Planet      struct {
		ID           string      `json:"id"`
		Name         string      `json:"name"`
		Resource     string      `json:"resource"`
		PlanetType   PlanetTypes `json:"planet_type"`
		Distance     int         `json:"distance"`
		OrbitalOrder int         `json:"orbital_order"`
		CreatedAt    time.Time   `json:"created_at"`
	}
)

const (
	Terrestrial PlanetTypes = "Terrestrial" // Spanish: terrestre
	GasGiant    PlanetTypes = "GasGiant"    // Spanish: gigante gaseoso
	IceGiant    PlanetTypes = "IceGiant"    // Spanish: gigante de hielo
	DwarfPlanet PlanetTypes = "DwarfPlanet" // Spanish: planeta enano
	Exoplanet   PlanetTypes = "Exoplanet"   // Spanish: exoplaneta
	aquatic     PlanetTypes = "aquatic"     // Spanish: acuático
	Desert      PlanetTypes = "Desert"      // Spanish: desierto
	Unknown     PlanetTypes = "Unknown"     // Spanish: desconocido
)
