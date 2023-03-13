CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS projects
(
    id           uuid    DEFAULT uuid_generate_v4(),
    project_name text NOT NULL UNIQUE,
    active       boolean DEFAULT false,
    CONSTRAINT pk_projects PRIMARY KEY (id)
);
