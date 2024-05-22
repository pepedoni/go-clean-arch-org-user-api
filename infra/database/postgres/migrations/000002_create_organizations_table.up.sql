CREATE TABLE "organizations" (
  "id" uuid PRIMARY KEY,
  "name" text NOT NULL,
  "document" text NOT NULL
);

COMMENT ON TABLE "organizations" IS 'Registro de organizações';