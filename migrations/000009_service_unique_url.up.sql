ALTER TABLE services
    ADD CONSTRAINT unique_service_url UNIQUE (url);