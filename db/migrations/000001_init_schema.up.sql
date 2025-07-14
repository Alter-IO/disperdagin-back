-- ========== INDEPENDENT BASE TABLES ==========

CREATE TABLE "roles" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR UNIQUE NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "subdistricts" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "commodity_types" (
  "id" VARCHAR PRIMARY KEY,
  "description" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "photo_categories" (
  "id" VARCHAR PRIMARY KEY,
  "category" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "markets" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

-- ========== DEPENDENT TABLES (LEVEL 1) ==========

CREATE TABLE "users" (
  "id" VARCHAR PRIMARY KEY,
  "role_id" VARCHAR NOT NULL,
  "username" VARCHAR UNIQUE NOT NULL,
  "password" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "deleted_at" TIMESTAMPTZ,
  CONSTRAINT fk_role FOREIGN KEY ("role_id") REFERENCES "roles" ("id")
);

CREATE TABLE "villages" (
  "id" VARCHAR PRIMARY KEY,
  "subdistrict_id" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ,
  CONSTRAINT fk_subdistrict FOREIGN KEY ("subdistrict_id") REFERENCES "subdistricts" ("id")
);

CREATE TABLE "commodities" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "unit" VARCHAR NOT NULL,
  "description" TEXT,
  "commodity_type_id" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ,
  CONSTRAINT fk_commodity_type FOREIGN KEY ("commodity_type_id") REFERENCES "commodity_types" ("id")
);

CREATE TABLE "daily_commodities" (
  "id" VARCHAR PRIMARY KEY,
  "commodities" JSONB, -- [{"id": "1", "price": 1000}]
  "publish_date" DATE NOT NULL
);

CREATE TABLE "market_fees" (
  "id" VARCHAR PRIMARY KEY,
  "market_id" VARCHAR NOT NULL,
  "num_permanent_kiosks" INTEGER NOT NULL,
  "num_non_permanent_kiosks" INTEGER NOT NULL,
  "permanent_kiosk_revenue" DECIMAL(15,2) NOT NULL,
  "non_permanent_kiosk_revenue" DECIMAL(15,2) NOT NULL,
  "collection_status" VARCHAR NOT NULL,
  "description" TEXT,
  "semester" VARCHAR NOT NULL,
  "year" INTEGER NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ,
  CONSTRAINT fk_market FOREIGN KEY ("market_id") REFERENCES "markets" ("id")
);

-- ========== TABLES WITHOUT DEPENDENCIES ==========

CREATE TABLE "news" (
  "id" VARCHAR PRIMARY KEY,
  "title" VARCHAR NOT NULL,
  "content" TEXT NOT NULL,
  "author" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "sectors" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "photos" (
  "id" VARCHAR PRIMARY KEY,
  "category_id" VARCHAR NOT NULL,
  "title" VARCHAR NOT NULL,
  "file_url" VARCHAR,
  "description" TEXT,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "ikms" (
  "id" VARCHAR PRIMARY KEY,
  "description" TEXT NOT NULL,
  "village_id" VARCHAR NOT NULL,
  "business_type" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "public_information" (
  "id" VARCHAR PRIMARY KEY,
  "document_name" VARCHAR NOT NULL,
  "file_url" VARCHAR NOT NULL,
  "public_info_type" VARCHAR NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "legal_documents" (
  "id" VARCHAR PRIMARY KEY,
  "document_name" VARCHAR NOT NULL,
  "file_url" VARCHAR NOT NULL,
  "document_type" VARCHAR NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "ikm_types" (
  "id" VARCHAR PRIMARY KEY,
  "document_name" VARCHAR NOT NULL,
  "file_url" VARCHAR NOT NULL,
  "public_info_type" VARCHAR NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "public_information_types" (
  "id" VARCHAR PRIMARY KEY,
  "description" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "legal_document_types" (
  "id" VARCHAR PRIMARY KEY,
  "description" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "employees" (
  "id" VARCHAR PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "position" VARCHAR NOT NULL,
  "address" TEXT,
  "employee_id" VARCHAR,
  "birthplace" VARCHAR,
  "birthdate" DATE,
  "photo" VARCHAR,
  "status" INTEGER,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "greetings" (
  "id" VARCHAR PRIMARY KEY,
  "message" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "videos" (
  "id" VARCHAR PRIMARY KEY,
  "title" VARCHAR NOT NULL,
  "link" VARCHAR NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);

CREATE TABLE "vision_mission" (
  "id" VARCHAR PRIMARY KEY,
  "vision" TEXT NOT NULL,
  "mission" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ,
  "author" VARCHAR NOT NULL,
  "deleted_at" TIMESTAMPTZ
);