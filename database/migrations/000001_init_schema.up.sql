CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_email_verified" boolean NOT NULL DEFAULT false,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "is_admin" boolean not null DEFAULT false,
  "is_active" boolean not null DEFAULT true,
  "deactivated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "is_deleted" boolean not null DEFAULT false,
  "deleted_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Verify_Emails" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes'),

  FOREIGN KEY ("user_id") REFERENCES "Users" ("id")
);

CREATE TABLE "Sessions" (
  "id" uuid PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),

  FOREIGN KEY ("user_id") REFERENCES "Users" ("id")
);

CREATE TABLE "Books" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "tags_array" integer[],
  "price" int NOT NULL,
  "description" varchar not null,
  "download_link" varchar not null,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Tags" (
  "id" int unique NOT NULL,
  "tag_name" varchar NOT NULL
);

CREATE TABLE "Reviews" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "book_id" bigserial NOT NULL,
  "rating" int NOT NULL,
  "comment" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),

  FOREIGN KEY ("user_id") REFERENCES "Users" ("id"),
  FOREIGN KEY ("book_id") REFERENCES "Books" ("id"),
  constraint "reviews_user_book_key" unique ("user_id", "book_id")
);

CREATE TABLE "Carts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "book_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),

  FOREIGN KEY ("user_id") REFERENCES "Users" ("id"),
  FOREIGN KEY ("book_id") REFERENCES "Books" ("id"),
  CONSTRAINT "carts_user_book_key" UNIQUE ("user_id", "book_id")
);

CREATE TABLE "Orders" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "razorpay_order_id" varchar unique not null,
  "user_id" bigserial NOT NULL,
  "total_money" float NOT NULL,
  "status" varchar NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),

  FOREIGN KEY ("user_id") REFERENCES "Users" ("id"),
  CONSTRAINT "orders_razor_user_key" UNIQUE ("razorpay_order_id", "user_id")
);

CREATE TABLE "Purchases" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "book_id" bigserial NOT NULL,
  "order_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),

  FOREIGN KEY ("book_id") REFERENCES "Books" ("id"),
  FOREIGN KEY ("order_id") REFERENCES "Orders" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "Users" ("id"),
  CONSTRAINT "order_lines_orderid_book_key" UNIQUE ("order_id", "book_id")
);

-- indexes
-- TODO: Create indexes wherever possible to speed things up

-- extension
create extension fuzzystrmatch;
