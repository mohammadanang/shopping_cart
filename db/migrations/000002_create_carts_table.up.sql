CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "product_name" varchar(180) NOT NULL,
  "qty" int NOT NULL DEFAULT 0,
  "buyer" varchar(120) NOT NULL,
  "price" float NOT NULL,
  "order_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "carts" ("buyer");

ALTER TABLE "carts" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");