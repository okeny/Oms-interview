-- +migrate Up
CREATE TABLE IF NOT EXISTS building (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP without TIME ZONE NOT NULL,
    updated_at TIMESTAMP without TIME ZONE NOT NULL
);

INSERT INTO building (name, address, created_at, updated_at) VALUES 
('Sunset Towers', '123 Sunshine Ave, Springfield', NOW(), NOW()),
('Maple Residency', '456 Maple Street, Riverdale', NOW(), NOW()),
('Skyline Heights', '789 Sky Drive, Metropolis', NOW(), NOW()),
('Pine View Apartments', '321 Pine Rd, Gotham', NOW(), NOW()),
('Ocean Breeze Complex', '654 Coastal Blvd, Atlantis', NOW(), NOW()),
('Greenfield Estate', '987 Greenfield Lane, Star City', NOW(), NOW()),
('Hilltop Haven', '159 Hilltop Ave, Central City', NOW(), NOW()),
('Lakeside Residency', '753 Lakeside Dr, Emerald City', NOW(), NOW()),
('Golden Gate Plaza', '852 Bridge St, San Francisco', NOW(), NOW()),
('Crystal Towers', '147 Crystal Rd, Angel Grove', NOW(), NOW());


-- +migrate Down
DROP TABLE IF EXISTS building;
