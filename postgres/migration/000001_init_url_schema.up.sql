CREATE TABLE "urls" (
  "id" bigint PRIMARY KEY,
  "url" text UNIQUE NOT NULL,
  "hash" varchar(6) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "urls" ("url");

CREATE INDEX ON "urls" ("hash");
