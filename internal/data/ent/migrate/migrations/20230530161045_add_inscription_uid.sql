-- Modify "collections" table
ALTER TABLE "collections" ADD COLUMN "inscription_uid" character varying NOT NULL;
-- Create index "collections_inscription_id_key" to table: "collections"
CREATE UNIQUE INDEX "collections_inscription_id_key" ON "collections" ("inscription_id");
-- Create index "collections_inscription_uid_key" to table: "collections"
CREATE UNIQUE INDEX "collections_inscription_uid_key" ON "collections" ("inscription_uid");
-- Modify "tokens" table
ALTER TABLE "tokens" ADD COLUMN "inscription_uid" character varying NOT NULL;
-- Create index "tokens_inscription_id_key" to table: "tokens"
CREATE UNIQUE INDEX "tokens_inscription_id_key" ON "tokens" ("inscription_id");
-- Create index "tokens_inscription_uid_key" to table: "tokens"
CREATE UNIQUE INDEX "tokens_inscription_uid_key" ON "tokens" ("inscription_uid");
