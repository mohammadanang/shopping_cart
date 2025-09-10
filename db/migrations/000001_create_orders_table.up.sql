CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "order_number" varchar(50) NOT NULL,
  "discount" float NOT NULL,
  "status" varchar(30) NOT NULL DEFAULT 'pending',
  "total" float NOT NULL,
  "buyer" varchar(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "orders" ("order_number");

COMMENT ON COLUMN "orders"."status" IS 'pending, completed';