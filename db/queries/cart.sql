-- name: AddCart :one
INSERT INTO carts (product_name, qty, price, order_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListCarts :many
SELECT *
FROM carts
WHERE order_id = $1
ORDER BY updated_at DESC;

-- name: GetCart :one
SELECT * FROM carts
WHERE id = $1 LIMIT 1;

-- name: EditCart :one
UPDATE carts
SET qty = $2,
  price = $3,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: RemoveCart :exec
DELETE FROM carts
WHERE id = $1;

-- name: RemoveCartsByOrder :exec
DELETE FROM carts
WHERE order_id = $1;