CREATE TABLE "organizations_users" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "organization_id" uuid NOT NULL
);

CREATE UNIQUE INDEX ON "organizations_users" ("user_id", "organization_id");

COMMENT ON TABLE "organizations_users" IS 'Relacionamento de usuários e organizações';

ALTER TABLE "organizations_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "organizations_users" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");