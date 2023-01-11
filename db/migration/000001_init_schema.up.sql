CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "tokenRefresh" varchar NOT NULL,
  "status" boolean NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT(now())
);

CREATE TABLE "contents" (
  "id" integer PRIMARY KEY,
  "fileid" varchar NOT NULL,
  "name" varchar NOT NULL,
  "size" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT(now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "contents" ("fileid");
