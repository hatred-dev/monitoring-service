ALTER TABLE ip_address
    DROP CONSTRAINT fk_ip_address_projects;

ALTER TABLE ip_address
    ADD CONSTRAINT fk_ip_address_projects FOREIGN KEY (project_id) REFERENCES projects (id);
