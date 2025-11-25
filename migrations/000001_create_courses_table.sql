-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL CHECK (length(title) > 0 AND length(title) <= 255),
    description TEXT,
    author TEXT NOT NULL CHECK (length(author) > 0),
    created_at TIMESTAMPTZ NOT NULL,      -- ← без DEFAULT
    updated_at TIMESTAMPTZ NOT NULL       -- ← без DEFAULT
);

CREATE INDEX idx_courses_author ON courses (author);
CREATE INDEX idx_courses_created_at ON courses (created_at DESC);

-- +goose Down
DROP INDEX IF EXISTS idx_courses_author;
DROP INDEX IF EXISTS idx_courses_created_at;
DROP TABLE IF EXISTS courses;
DROP EXTENSION IF EXISTS "uuid-ossp";