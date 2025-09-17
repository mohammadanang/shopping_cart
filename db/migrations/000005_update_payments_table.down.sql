ALTER TABLE payments
DROP COLUMN "status",
DROP COLUMN "gateway",
DROP COLUMN "gateway_ref_id",
DROP COLUMN "raw_response";