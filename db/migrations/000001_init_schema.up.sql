-- ========== INDEPENDENT BASE TABLES ==========

-- Create roles table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'roles') THEN
        CREATE TABLE "roles" (
          "id" VARCHAR PRIMARY KEY,
          "name" VARCHAR UNIQUE NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create subdistricts table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'subdistricts') THEN
        CREATE TABLE "subdistricts" (
          "id" VARCHAR PRIMARY KEY,
          "name" VARCHAR NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create commodity_types table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'commodity_types') THEN
        CREATE TABLE "commodity_types" (
          "id" VARCHAR PRIMARY KEY,
          "description" TEXT NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create photo_categories table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'photo_categories') THEN
        CREATE TABLE "photo_categories" (
          "id" VARCHAR PRIMARY KEY,
          "category" VARCHAR NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create markets table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'markets') THEN
        CREATE TABLE "markets" (
          "id" VARCHAR PRIMARY KEY,
          "name" VARCHAR NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- ========== DEPENDENT TABLES (LEVEL 1) ==========

-- Create users table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'users') THEN
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
    END IF;
END
$$;

-- Create villages table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'villages') THEN
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
    END IF;
END
$$;

-- Create commodities table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'commodities') THEN
        CREATE TABLE "commodities" (
          "id" VARCHAR PRIMARY KEY,
          "name" VARCHAR NOT NULL,
          "price" DECIMAL(15,2) NOT NULL,
          "unit" VARCHAR NOT NULL,
          "publish_date" DATE NOT NULL,
          "description" TEXT,
          "commodity_type_id" VARCHAR NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ,
          CONSTRAINT fk_commodity_type FOREIGN KEY ("commodity_type_id") REFERENCES "commodity_types" ("id")
        );
    END IF;
END
$$;

-- Create market_fees table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'market_fees') THEN
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
    END IF;
END
$$;

-- ========== TABLES WITHOUT DEPENDENCIES ==========

-- Create news table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'news') THEN
        CREATE TABLE "news" (
          "id" VARCHAR PRIMARY KEY,
          "title" VARCHAR NOT NULL,
          "content" TEXT NOT NULL,
          "author" VARCHAR NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create sectors table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'sectors') THEN
        CREATE TABLE "sectors" (
          "id" VARCHAR PRIMARY KEY,
          "name" VARCHAR NOT NULL,
          "description" TEXT,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create photos table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'photos') THEN
        CREATE TABLE "photos" (
          "id" VARCHAR PRIMARY KEY,
          "category_id" VARCHAR NOT NULL,
          "title" VARCHAR NOT NULL,
          "file" VARCHAR,
          "description" TEXT,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create ikms table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'ikms') THEN
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
    END IF;
END
$$;

-- Create public_information table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'public_information') THEN
        CREATE TABLE "public_information" (
          "id" VARCHAR PRIMARY KEY,
          "document_name" VARCHAR NOT NULL,
          "file_name" VARCHAR NOT NULL,
          "public_info_type" VARCHAR NOT NULL,
          "description" TEXT,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create legal_documents table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'legal_documents') THEN
        CREATE TABLE "legal_documents" (
          "id" VARCHAR PRIMARY KEY,
          "document_name" VARCHAR NOT NULL,
          "file_name" VARCHAR NOT NULL,
          "document_type" VARCHAR NOT NULL,
          "description" TEXT,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create ikm_types table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'ikm_types') THEN
        CREATE TABLE "ikm_types" (
          "id" VARCHAR PRIMARY KEY,
          "document_name" VARCHAR NOT NULL,
          "file_name" VARCHAR NOT NULL,
          "public_info_type" VARCHAR NOT NULL,
          "description" TEXT,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create public_information_types table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'public_information_types') THEN
        CREATE TABLE "public_information_types" (
          "id" VARCHAR PRIMARY KEY,
          "description" TEXT NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create legal_document_types table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'legal_document_types') THEN
        CREATE TABLE "legal_document_types" (
          "id" VARCHAR PRIMARY KEY,
          "description" TEXT NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create employees table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'employees') THEN
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
    END IF;
END
$$;

-- Create greetings table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'greetings') THEN
        CREATE TABLE "greetings" (
          "id" VARCHAR PRIMARY KEY,
          "message" TEXT NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;

-- Create videos table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'videos') THEN
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
    END IF;
END
$$;

-- Create vision_mission table if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'vision_mission') THEN
        CREATE TABLE "vision_mission" (
          "id" VARCHAR PRIMARY KEY,
          "vision" TEXT NOT NULL,
          "mission" TEXT NOT NULL,
          "created_at" TIMESTAMPTZ NOT NULL,
          "updated_at" TIMESTAMPTZ,
          "author" VARCHAR NOT NULL,
          "deleted_at" TIMESTAMPTZ
        );
    END IF;
END
$$;