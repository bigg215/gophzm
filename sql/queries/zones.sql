-- name: AddZip :one
INSERT INTO zones (
	id,
	createdat,
	updatedat,
	zipcode,
	zone,
	temprange,
	zonetitle
 )
 VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
 )
 RETURNING *;

 -- name: GetZipZone :one
 SELECT *
 FROM zones
 WHERE zipcode = $1;