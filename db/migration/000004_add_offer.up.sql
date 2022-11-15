CREATE TABLE "offers" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "from_currency" varchar NOT NULL,
    "to_currency" varchar NOT NULL,
    "amount" float(32) NOT NULL,
    "rate" float(32) NOT NULL,
    "status" varchar NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "offers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;