CREATE TYPE job_status AS ENUM ('created', 'executing', 'done', 'failed');
CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    file_id INT NOT NULL,
    status job_status NOT NULL DEFAULT 'created',
    query VARCHAR(512),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_change TIMESTAMP,
);