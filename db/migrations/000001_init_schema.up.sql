CREATE TABLE "roles" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR UNIQUE NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "users" (
  "id" VARCHAR PRIMARY KEY,
  "role_id" VARCHAR NOT NULL,
  "username" VARCHAR UNIQUE NOT NULL,
  "password" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "deleted_at" TIMESTAMPTZ
);