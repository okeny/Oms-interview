-- +migrate Up

CREATE TABLE IF NOT EXISTS apartment (
    id SERIAL PRIMARY KEY,
    building_id INTEGER NOT NULL REFERENCES building(id) ON DELETE CASCADE,
    number VARCHAR(50) NOT NULL,
    floor INTEGER NOT NULL,
    sq_meters INTEGER NOT NULL,
    created_at TIMESTAMP without TIME ZONE NOT NULL,
    updated_at TIMESTAMP without TIME ZONE NOT NULL
);

INSERT INTO apartment (building_id, number, floor, sq_meters, created_at, updated_at) VALUES
(1, 'A101', 1, 65,NOW(), NOW()),
(1, 'A102', 1, 70, NOW(), NOW()),
(2, 'B201', 2, 80,NOW(), NOW()),
(2, 'B202', 2, 75,NOW(), NOW()),
(3, 'C301', 3, 90,NOW(), NOW()),
(3, 'C302', 3, 85,NOW(), NOW()),
(4, 'D401', 4, 100,NOW(), NOW()),
(4, 'D402', 4, 95,NOW(), NOW()),
(5, 'E501', 5, 110,NOW(), NOW()),
(5, 'E502', 5, 105,NOW(), NOW());

-- +migrate Down
DROP TABLE IF EXISTS apartment;
