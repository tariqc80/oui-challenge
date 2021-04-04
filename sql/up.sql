CREATE TABLE IF NOT EXISTS sets (
    id SERIAL PRIMARY KEY,
    members integer[]
);

ALTER TABLE sets
    ADD COLUMN hash text DEFAULT null;
