-- +goose Up
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE subscriptions(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    service_name varchar(1024) NOT NULL,
    user_id UUID NOT NULL,
    price integer NOT NULL,
    start_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    end_date TIMESTAMP WITHOUT TIME ZONE
);

create index service_name_idx on public.subscriptions using gin(lower(service_name) gin_trgm_ops);
create index user_id_idx on public.subscriptions (user_id);

-- +goose Down
DROP INDEX IF EXISTS service_name_idx;
DROP INDEX IF EXISTS user_id_idx;
DROP TABLE IF EXISTS subscriptions;