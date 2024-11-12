package domain

type (
	PlanetTypes string
	Planet      struct {
		ID         string
		Name       string
		Resource   string
		PlanetType PlanetTypes
		distance   int
		order      int
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
