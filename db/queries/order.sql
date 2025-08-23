-- name: AddOrder :one
INSERT INTO orders (order_number, discount, "status", total)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: PaginateOrders :many
SELECT * FROM orders
ORDER BY updated_at DESC
LIMIT $1
OFFSET $2;

-- name: PaginateOrdersWithParams :many
SELECT * FROM orders
WHERE order_number ILIKE $3
ORDER BY updated_at DESC
LIMIT $1
OFFSET $2;

-- name: CountOrdersByOrderNumber :one
SELECT COUNT(*) FROM orders
WHERE order_number ILIKE $1;

-- name: CountOrders :one
SELECT COUNT(*) FROM orders;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: EditOrder :one
UPDATE orders
SET discount = $2,
  "status" = $3,
  total = $4,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: RemoveOrder :exec
DELETE FROM orders
WHERE id = $1;