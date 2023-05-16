CREATE TYPE account_types AS ENUM (
  'ASSET',
  'LIABILITY'
);

CREATE TYPE product_types AS ENUM (
  'CREDIT_CARD',
  'CHEQUING',
  'SAVINGS'
);

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "accounts" ADD COLUMN account_type account_types;
ALTER TABLE "accounts" ADD COLUMN product_type product_types;
