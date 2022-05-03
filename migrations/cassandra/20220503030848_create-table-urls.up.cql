CREATE TABLE short.urls
(
    id           VARCHAR PRIMARY KEY,
    hash         VARCHAR,
    original_url VARCHAR,
    created_at   DATE,
    updated_at   DATE,
    expired_at   DATE,
    user_id      VARCHAR
);

CREATE CUSTOM INDEX urls__index
    ON urls (hash) USING 'org.apache.cassandra.index.sasi.SASIIndex';
