CREATE TABLE "accesses" (
  "id" bigserial PRIMARY KEY,
  "api_key" text NOT NULL,
  "refresh_token" text NOT NULL,
  "role" varchar(30) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accesses" ("api_key");

CREATE UNIQUE INDEX ON "accesses" ("refresh_token");

COMMENT ON COLUMN "accesses"."role" IS '`admin`, `guest`, `user`';