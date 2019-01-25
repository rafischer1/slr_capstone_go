CREATE TABLE admin
(
  id SERIAL NOT NULL PRIMARY KEY,
  username VARCHAR NOT NULL,
  password VARCHAR NOT NULL
);

INSERT INTO admin
  (id, username, password)
VALUES
  (1, 'gmri', 'slrmaine');

DROP TABLE admin;

-- dev: psql capstonedb