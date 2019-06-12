CREATE TABLE IF NOT EXISTS dely
(
    id INTEGER PRIMARY KEY,
    lat FLOAT, 
    lon FLOAT,
    geom GEOMETRY(POINT, 4326), 
    fuel VARCHAR(10), 
    name VARCHAR(255),
    year INTEGER,
    engine_capacity FLOAT,
    engine_power INTEGER,
    transmission VARCHAR(255),
    equipment VARCHAR(255), 
    img VARCHAR(255),
    name_full VARCHAR(255),
    img_thumb VARCHAR(255)
);
CREATE INDEX ON dely(id);