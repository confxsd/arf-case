-- name: CreateOffer :one
INSERT INTO offers (
  user_id,
  from_currency,
  to_currency,
  rate,
  amount
) VALUES (
  $1, $2, $3, $4, $5 
) RETURNING *;


-- name: GetOffer :one
SELECT * FROM offers
WHERE id = $1 LIMIT 1;