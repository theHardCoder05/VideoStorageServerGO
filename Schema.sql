CREATE TABLE "user" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "tokenRefresh" varchar NOT NULL,
  "status" boolean NOT NULL,
  "created_at" datetime NOT NULL
);

CREATE TABLE "content" (
  "id" integer PRIMARY KEY,
  "fileid" varchar NOT NULL,
  "name" varchar NOT NULL,
  "size" bigint NOT NULL,
  "created_at" datetime NOT NULL
);

CREATE INDEX ON "user" ("username");

CREATE INDEX ON "content" ("fileid");
