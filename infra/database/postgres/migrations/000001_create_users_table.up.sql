CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "phone" text NOT NULL,
  "document" text NOT NULL
);

COMMENT ON TABLE "users" IS 'Registro de usuários';