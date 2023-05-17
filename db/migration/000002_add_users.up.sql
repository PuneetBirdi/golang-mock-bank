CREATE TYPE ACCOUNTS_ENUM AS ENUM (
  'ASSET',
  'LIABILITY'
);

CREATE TYPE PRODUCTS_ENUM AS ENUM (
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

ALTER TABLE "accounts" ADD COLUMN account_type ACCOUNTS_ENUM NOT NULL;
ALTER TABLE "accounts" ADD COLUMN product_type PRODUCTS_ENUM NOT NULL;
