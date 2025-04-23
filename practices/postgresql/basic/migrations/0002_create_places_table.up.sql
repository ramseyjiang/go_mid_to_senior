CREATE TABLE places
(
    id        SERIAL PRIMARY KEY,
    name      TEXT,
    longitude FLOAT,
    latitude  FLOAT,
    user_id   INT REFERENCES users (id) ON DELETE CASCADE
);
