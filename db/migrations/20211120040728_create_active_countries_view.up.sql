CREATE MATERIALIZED VIEW IF NOT EXISTS active_countries AS SELECT DISTINCT country FROM users;