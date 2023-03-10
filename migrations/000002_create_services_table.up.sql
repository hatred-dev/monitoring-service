CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE services
(
    id           uuid DEFAULT uuid_generate_v4(),
    project_name text NOT NULL,
    service_name text NOT NULL,
    url          text NOT NULL,
    CONSTRAINT pk_services PRIMARY KEY (id)
);

ALTER TABLE services
    ADD CONSTRAINT fk_services_projects FOREIGN KEY (project_name) REFERENCES projects (project_name);
