CREATE TABLE subscribers
(
  id SERIAL NOT NULL PRIMARY KEY,
  phone VARCHAR NOT NULL,
  location VARCHAR
);

INSERT INTO subscribers
  (id, phone, location)
VALUES
  (1, '3024232120', 'Cape Elizabeth'),
  (2, '2075556748', 'Portland');

ALTER TABLE subscribers ADD UNIQUE (phone);
ALTER SEQUENCE subscribers_id_seq RESTART WITH 3;

DROP TABLE subscribers;