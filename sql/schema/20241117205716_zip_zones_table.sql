-- +goose Up
CREATE TABLE zones (
	id uuid NOT NULL PRIMARY KEY,
	createdat TIMESTAMP NOT NULL,
	updatedat TIMESTAMP NOT NULL,
	zipcode VARCHAR NOT NULL, 
	zone VARCHAR NOT NULL,
	temprange VARCHAR NOT NULL,
	zonetitle VARCHAR NOT NULL 
);

CREATE UNIQUE INDEX zip_uidx ON zones(zipcode);

-- +goose Down

DROP INDEX IF EXISTS zip_uidx;
DROP TABLE IF EXISTS zones;
