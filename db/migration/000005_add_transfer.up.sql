CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_wallet_id" bigint NOT NULL,
  "to_wallet_id" bigint NOT NULL,
  "amount" float(32) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_wallet_id") REFERENCES "wallets" ("id") ON DELETE CASCADE;

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_wallet_id") REFERENCES "wallets" ("id") ON DELETE CASCADE;