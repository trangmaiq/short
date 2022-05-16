CREATE TABLE short.urls
(
    hash         VARCHAR PRIMARY KEY ,
    original_url VARCHAR,
    created_at   DATE,
    updated_at   DATE,
    expired_at   DATE,
    user_id      VARCHAR
);
