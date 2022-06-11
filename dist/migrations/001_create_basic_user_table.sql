CREATE TABLE "user" (
    id uuid not null primary key,
    email text not null,
    password text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone
);

CREATE UNIQUE INDEX IF NOT EXISTS "user_idx_email_uq" ON "user" ("email");

---- create above / drop below ----

DROP INDEX IF EXISTS "user_idx_email";
DROP TABLE IF EXISTS "user";
