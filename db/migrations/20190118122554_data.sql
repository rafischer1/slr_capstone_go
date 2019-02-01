DROP TABLE data;

CREATE TABLE data
(
  id SERIAL NOT NULL PRIMARY KEY,
  msg VARCHAR NOT NULL,
  windmph REAL,
  winddir VARCHAR,
  sealevelft REAL,
  category VARCHAR,
  createdat timestamp default current_timestamp
);

INSERT INTO data
  (id, msg, windmph, winddir, sealevelft, category)
VALUES
  (1, 'Flooding expected Portland
Waterfront this afternoon. Tide level 11.6ft
with 20-30mph winds will result in Minor flooding along piers and streets around 2:30pm. Please see: https:
//slr-maine.herokuapp.com/ for more information', 20, 'NE', 11.8, 'Minor'),
(2, 'Major Flooding expected Portland
Waterfront overnight. Tide level 11.8+ft
with 30-35mph winds and period of heavy rain will result in Major flooding along piers and streets. Please see: https:
//twitter.com/cityportland for more information', 35, 'SW', 11.8, 'Major'
),
(3, '"Splash Over" event expected Portland
Waterfront this morning. Tide level 11.6ft
with 10-15mph winds and light rain will result in splashing on waterfront piers', 15, 'NE', 11.6, 'Splash Over'
),
(4, 'Minor flooding event expected Portland
Harbor today. Tide level 11.8ft
with 10-15mph winds will result in "King Tide" flooding along piers and waterfront streets', 10, 'S', 11.8, 'Minor'
);

ALTER SEQUENCE data_id_seq RESTART WITH 5;

-- dev: psql capstonedb
