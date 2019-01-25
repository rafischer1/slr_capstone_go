CREATE TABLE data
(
  id SERIAL NOT NULL PRIMARY KEY,
  msg VARCHAR NOT NULL,
  windmph REAL,
  winddir VARCHAR,
  sealevelft REAL,
  createdat timestamp default current_timestamp
);

ALTER SEQUENCE data_id_seq RESTART WITH 1;

DROP TABLE data;

-- dev: psql capstonedb
