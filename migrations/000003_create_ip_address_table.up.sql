CREATE TABLE IF NOT EXISTS ip_address
(
    id           uuid DEFAULT uuid_generate_v4(),
    project_name text NOT NULL,
    ip           text NOT NULL,
    CONSTRAINT pk_ip_address PRIMARY KEY (id)
);

ALTER TABLE ip_address
    ADD CONSTRAINT fk_ip_address_projects FOREIGN KEY (project_name) REFERENCES projects (project_name);
