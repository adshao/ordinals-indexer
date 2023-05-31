-- Drop index "token_inscription_id" from table: "tokens"
DROP INDEX "token_inscription_id";
-- Create index "token_inscription_id" to table: "tokens"
CREATE UNIQUE INDEX "token_inscription_id" ON "tokens" ("inscription_id");
