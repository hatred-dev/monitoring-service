ALTER TABLE ip_address
    DROP CONSTRAINT fk_ip_address_projects;

ALTER TABLE ip_address
    ADD CONSTRAINT fk_ip_address_projects FOREIGN KEY (project_name) REFERENCES projects (project_name) ON DELETE CASCADE;
