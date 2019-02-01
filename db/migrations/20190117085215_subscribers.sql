DROP TABLE subscribers;

CREATE TABLE subscribers
(
  id SERIAL NOT NULL PRIMARY KEY,
  phone VARCHAR NOT NULL,
  location VARCHAR
);

ALTER TABLE subscribers ADD UNIQUE (phone);

INSERT INTO subscribers
  (id, phone, location)
VALUES
  (1, '3024232120', 'Cape Elizabeth'),
  (2, '2075556748', 'Portland'),
  (3, '2075556749', 'Other'),
  (4, '2075556755', 'Portland'),
(5, '2075556741', 'Portland'),
(6, '2075556742', 'South Portland'),
(7, '2075556743', 'South Portland'),
(8, '2075556744', 'Falmouth'),
(9, '2075556745', 'Westbrook'),
(10, '2075556746', 'Westbrook'),
(11, '2075556747', 'Westbrook'),
(12, '2085556748', 'Westbrook'),
(13, '2078556748', 'Cape Elizabeth'),
(14, '2075856748', 'Cape Elizabeth'),
(15, '2075586748', 'Cape Elizabeth'),
(16, '2075558748', 'Cape Elizabeth'),
(17, '2075556848', 'South Portland'),
(18, '1075556748', 'South Portland'),
(19, '2175556748', 'South Portland'),
(20, '2015556748', 'Falmouth'),
(21, '2071556748', 'Falmouth'),
(22, '2075156748', 'Falmouth'),
(23, '2075516748', 'Portland'),
(24, '2075551748', 'Portland'),
(25, '2075556148', 'Portland'),
(26, '2275556748', 'Portland'),
(27, '2025556748', 'Portland'),
(28, '3275556749', 'Other'),
(29, '3275556755', 'Portland'),
(30, '3275556741', 'Portland'),
(31, '3275556742', 'South Portland'),
(32, '3275556743', 'South Portland'),
(33, '3275556744', 'Falmouth'),
(34, '3275556745', 'Westbrook'),
(35, '3275556746', 'Westbrook'),
(36, '3275556747', 'Westbrook'),
(37, '3285556748', 'Westbrook'),
(38, '3278556748', 'Cape Elizabeth'),
(39, '3275856748', 'Cape Elizabeth'),
(40, '3275586748', 'Cape Elizabeth'),
(41, '3275558748', 'Cape Elizabeth'),
(42, '3275556848', 'South Portland'),
(43, '4075556748', 'South Portland'),
(44, '5175556748', 'South Portland'),
(45, '9215556748', 'Falmouth'),
(46, '3271556748', 'Falmouth'),
(47, '3275156748', 'Falmouth'),
(48, '3275516748', 'Portland'),
(49, '3275551748', 'Portland'),
(50, '3275556148', 'Portland');

ALTER SEQUENCE subscribers_id_seq RESTART WITH 51;

