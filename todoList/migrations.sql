CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE
);