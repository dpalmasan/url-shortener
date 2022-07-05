CREATE DATABASE IF NOT EXISTS url_shortener;

USE url_shortener;

CREATE TABLE IF NOT EXISTS urls (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    shorten_url VARCHAR(255) NOT NULL
    url_id BIGINT NOT NULL
);

CREATE INDEX short_url ON urls(shorten_url);

/* For demo purposes */
ALTER TABLE urls AUTO_INCREMENT=18495736894938;