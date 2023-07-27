-- Modify "collections" table
ALTER TABLE "collections" ADD COLUMN "sig" jsonb NOT NULL;
-- Modify "tokens" table
ALTER TABLE "tokens" ADD COLUMN "sig" jsonb NOT NULL, ADD COLUMN "sig_uid" character varying NOT NULL;
