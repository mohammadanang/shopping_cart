CREATE TABLE "payments" (
  "id" bigserial PRIMARY KEY,
  "payment_number" varchar(50) NOT NULL,
  "order_id" bigint NOT NULL,
  "method" varchar(30) NOT NULL,
  "total" float NOT NULL,
  "paid_at" timestamptz NULL DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "payments" ("payment_number");

COMMENT ON COLUMN "payments"."method" IS 'bank, e-wallet, etc';

ALTER TABLE "payments" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");