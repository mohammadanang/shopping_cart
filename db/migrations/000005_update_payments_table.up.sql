ALTER TABLE payments
ADD COLUMN "status" varchar(30) NOT NULL DEFAULT 'pending',
ADD COLUMN gateway varchar(30),
ADD COLUMN gateway_ref_id varchar(200),
ADD COLUMN raw_response jsonb;

COMMENT ON COLUMN "payments"."status" IS 'pending, success, failed';

COMMENT ON COLUMN "payments"."gateway" IS 'xendit, stripe, midtrans';