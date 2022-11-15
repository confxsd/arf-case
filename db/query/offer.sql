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
