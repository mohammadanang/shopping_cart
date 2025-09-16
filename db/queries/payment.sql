-- name: AddPayment :one
INSERT INTO payments (payment_number, order_id, method, total)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: PaginatePayments :many
SELECT * FROM payments
ORDER BY updated_at DESC
LIMIT $1
OFFSET $2;

-- name: PaginatePaymentsWithParams :many
SELECT * FROM payments
WHERE payment_number ILIKE $3
ORDER BY updated_at DESC
LIMIT $1
OFFSET $2;

-- name: CountPayments :one
SELECT COUNT(*) FROM payments;

-- name: CountPaymentsByPaymentNumber :one
SELECT COUNT(*) FROM payments
WHERE payment_number ILIKE $1;

-- name: GetPayment :one
SELECT * FROM payments
WHERE payment_number = $1 LIMIT 1;

-- name: GetPaymentByOrder :one
SELECT * FROM payments
WHERE order_id = $1 LIMIT 1;

-- name: EditPayment :one
UPDATE payments
SET method = $2,
  total = $3,
  paid_at = $4,
  updated_at = now()
WHERE payment_number = $1
RETURNING *;