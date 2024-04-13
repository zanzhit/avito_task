CREATE TABLE users (
    id       SERIAL       PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash         VARCHAR(255) NOT NULL,
    role     VARCHAR(255)     NOT NULL
);

CREATE TABLE banners (
    id         SERIAL PRIMARY KEY,
    content    JSONB,
    is_active  BOOLEAN,
    feature_id INTEGER,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tags (
    tag_id INTEGER PRIMARY KEY
);

CREATE TABLE features (
    feature_id SERIAL PRIMARY KEY
);

CREATE TABLE banner_tags (
    id SERIAL NOT NULL UNIQUE,
    banner_id INT REFERENCES banners ON DELETE CASCADE NOT NULL,
    tags_id INT REFERENCES tags ON DELETE CASCADE NOT NULL
);