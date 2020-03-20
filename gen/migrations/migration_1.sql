-- +migrate Up
CREATE TYPE task_status AS ENUM ('running', 'created', 'finished');
CREATE TABLE tasks(
    uuid UUID PRIMARY KEY,
    status task_status DEFAULT 'created',
    created_at  timestamptz DEFAULT now()
);
-- +migrate Down
DROP TABLE tasks;
DROP TYPE task_status;