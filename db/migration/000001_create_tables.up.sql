CREATE TABLE "company" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "description" text NOT NULL, 
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "branch" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "name" text NOT NULL,
  "description" text NOT NULL, 
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "area" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "department" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "role" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "employee" (
  "id" uuid PRIMARY KEY,
  "number" bigint NOT NULL,
  "name" text NOT NULL,
  "surname" text NOT NULL,
  "birthdate" text NOT NULL,
  "dni" text NOT NULL,
  "cuil" text NOT NULL,
  "marital_status" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "contract" (
  "id" bigserial PRIMARY KEY,
  "type" text NOT NULL,
  "start_date" text NOT NULL,
  "employee_id" uuid NOT NULL,
  "company_id" bigint NOT NULL,
  "branch_id" bigint NOT NULL,
  "area_id" bigint NOT NULL,
  "department_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "paycheck" (
  "id" bigserial PRIMARY KEY,
  "type" text NOT NULL,
  "filename" text NOT NULL,
  "description" text NOT NULL,
  "folder" text NOT NULL,
  "path" text NOT NULL,
  "read" bool NOT NULL,
  "signed" bool NOT NULL,
  "employee_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "updated_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE INDEX ON "company" ("id");

CREATE INDEX ON "company" ("name");

CREATE INDEX ON "branch" ("id");

CREATE INDEX ON "branch" ("name");

CREATE INDEX ON "area" ("id");

CREATE INDEX ON "area" ("name");

CREATE INDEX ON "department" ("id");

CREATE INDEX ON "department" ("name");

CREATE INDEX ON "role" ("id");

CREATE INDEX ON "role" ("name");

CREATE INDEX ON "employee" ("id");

CREATE INDEX ON "employee" ("number");

CREATE INDEX ON "employee" ("name");

CREATE INDEX ON "employee" ("cuil");

CREATE INDEX ON "employee" ("dni");

CREATE INDEX ON "contract" ("id");

CREATE INDEX ON "contract" ("employee_id");

CREATE INDEX ON "contract" ("type");

CREATE INDEX ON "contract" ("start_date");

CREATE INDEX ON "paycheck" ("id");

CREATE INDEX ON "paycheck" ("employee_id");

CREATE INDEX ON "paycheck" ("read");

CREATE INDEX ON "paycheck" ("signed");

ALTER TABLE "branch" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "contract" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");

ALTER TABLE "contract" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "contract" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "contract" ADD FOREIGN KEY ("area_id") REFERENCES "area" ("id");

ALTER TABLE "contract" ADD FOREIGN KEY ("department_id") REFERENCES "department" ("id");

ALTER TABLE "contract" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "paycheck" ADD FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");
