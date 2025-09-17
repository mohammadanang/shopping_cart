-- name: AddPayment :one
INSERT INTO payments (payment_number, order_id, method, total, "status", gateway)
VALUES ($1, $2, $3, $4, $5, $6)
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
  gateway = $4,
  "status" = $5,
  gateway_ref_id = $6,
  raw_response = $7,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: ApprovePayment :one
UPDATE payments
SET paid_at = now(),
  total = $2,
  "status" = 'success',
  gateway_ref_id = $3,
  raw_response = $4,
  updated_at = now()
WHERE id = $1
RETURNING *;