ALTER TABLE "accounts" DROP COLUMN IF EXISTS "product_type";
ALTER TABLE "accounts" DROP COLUMN IF EXISTS "account_type";

CREATE TYPE PRODUCT_TYPE_ENUM AS ENUM (
	'ASSET',
	'LIABILITY'
);

DROP TYPE IF EXISTS PRODUCTS_ENUM;
DROP TYPE IF EXISTS ACCOUNTS_ENUM;
DROP TYPE IF EXISTS accounts_enum;
DROP TYPE IF EXISTS products_enum;

CREATE TABLE "products" (
	"id" bigserial PRIMARY KEY,
	"product_type" PRODUCT_TYPE_ENUM NOT NULL,
	"product_name" varchar NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO "products"(product_type, product_name) 
VALUES 
	('ASSET', 'SAVING'),
	('ASSET', 'CHEQUING'),
	('LIABILITY', 'CREDIT_CARD'),
	('LIABILITY', 'LINE_OF_CREDIT');

ALTER TABLE "accounts"
	ADD COLUMN "product_id" bigserial NOT NULL;

ALTER TABLE "accounts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
