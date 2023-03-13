ALTER TABLE services
    DROP CONSTRAINT fk_services_projects;
ALTER TABLE services
    ADD CONSTRAINT fk_services_projects FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE;
