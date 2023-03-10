CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS projects
(
    project_name text NOT NULL,
    active       boolean DEFAULT false,
    CONSTRAINT pk_projects PRIMARY KEY (project_name)
);