-- +goose Up
CREATE TABLE subscribers
(
  id SERIAL NOT NULL PRIMARY KEY,
  phone VARCHAR NOT NULL,
  location VARCHAR
);

ALTER TABLE subscribers ADD UNIQUE (phone);

-- +goose Down
DROP TABLE subscribers;