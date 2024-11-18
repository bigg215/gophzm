-- +goose Up
CREATE TABLE zones (
	id uuid NOT NULL PRIMARY KEY,
	createdat TIMESTAMP NOT NULL,
	updatedat TIMESTAMP NOT NULL,
	zipcode INT NOT NULL, 
	zone VARCHAR NOT NULL,
	temprange VARCHAR NOT NULL,
	zonetitle VARCHAR NOT NULL,
	year INT NOT NULL 
);

CREATE UNIQUE INDEX zip_uidx ON zones(zipcode, year);

-- +goose Down

DROP INDEX IF EXISTS zip_uidx;
DROP TABLE IF EXISTS zones;
