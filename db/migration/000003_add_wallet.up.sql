CREATE TABLE "wallets" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "balance" bigint NOT NULL DEFAULT 0,
    "currency" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "wallets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");