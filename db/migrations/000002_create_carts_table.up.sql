CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "product_name" varchar(180) NOT NULL,
  "qty" int NOT NULL DEFAULT 0,
  "price" float NOT NULL,
  "order_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "carts" ("product_name");

ALTER TABLE "carts" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");