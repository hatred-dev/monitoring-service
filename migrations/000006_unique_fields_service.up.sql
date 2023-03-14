ALTER TABLE services
    ADD CONSTRAINT unique_project_service UNIQUE (project_id, service_name);