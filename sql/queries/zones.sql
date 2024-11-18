-- name: AddZip :one
INSERT INTO zones (
	createdat,
	updatedat,
	zipcode,
	zone,
	temprange,
	zonetitle,
	year
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

-- name: GetZipsForZone :many
SELECT *
FROM zones
WHERE zone = $1;

-- name: DeleteAllZips :exec
DELETE FROM zones;
