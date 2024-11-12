-- Crear la tabla de usuarios
CREATE TABLE IF NOT EXISTS public.planets (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  resource VARCHAR(255) NOT NULL,
  planet_type VARCHAR(255) NOT NULL,
  distance int NOT NULL,
  orbital_order int NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Poblar datos iniciales
INSERT INTO planets (id, name, resource, planet_type, distance, orbital_order)
VALUES
  ('1', 'Mercury', 'Iron', 'Terrestrial', 36, 1),
  ('2', 'Venus', 'Copper', 'Terrestrial', 67, 2),
  ('3', 'Earth', 'Water', 'Terrestrial', 93, 3),
  ('4', 'Mars', 'Iron', 'Terrestrial', 141, 4),
  ('5', 'Jupiter', 'Helium', 'IceGiant', 483, 5),
  ('6', 'Saturn', 'Helium', 'IceGiant', 886, 6),
  ('7', 'Uranus', 'Helium', 'Desert', 1782, 7),
  ('8', 'Neptune', 'Helium', 'Desert', 2793, 8);