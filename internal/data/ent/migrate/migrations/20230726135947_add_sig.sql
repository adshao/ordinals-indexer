-- Modify "collections" table
ALTER TABLE "collections" ADD COLUMN "sig" jsonb NULL;
-- Modify "tokens" table
ALTER TABLE "tokens" ADD COLUMN "sig" jsonb NULL, ADD COLUMN "sig_uid" character varying NOT NULL DEFAULT '';
-- Create index "token_sig_uid" to table: "tokens"
CREATE INDEX "token_sig_uid" ON "tokens" ("sig_uid");
