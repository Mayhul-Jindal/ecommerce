CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "is_admin" boolean not null DEFAULT false,
  "is_active" boolean not null DEFAULT true,
  "deactivated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "is_deleted" boolean not null DEFAULT false,
  "deleted_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Sessions" (
  "id" uuid PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Addresses" (
  "user_id" bigserial NOT NULL,
  "address_line" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "country" varchar NOT NULL
);

CREATE TABLE "Books" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "tags_array" integer[],
  "price" int NOT NULL,
  "quantity" int NOT NULL DEFAULT 100,
  "description" varchar not null,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Tags" (
  "id" serial NOT NULL,
  "tag_name" varchar NOT NULL
);

CREATE TABLE "Reviews" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "book_id" bigserial NOT NULL,
  "rating" int NOT NULL,
  "comment" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Carts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Cart_Items" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "cart_id" bigserial NOT NULL,
  "book_id" bigserial NOT NULL,
  "quantity" int NOT NULL DEFAULT 1,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Orders" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "total_money" int NOT NULL,
  "order_status" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Order_Lines" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "book_id" bigserial NOT NULL,
  "order_id" bigserial NOT NULL,
  "quantity" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Users" ("email");

CREATE INDEX ON "Reviews" ("user_id", "book_id");

ALTER TABLE "Sessions" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Addresses" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Reviews" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Reviews" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");

ALTER TABLE "Carts" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Cart_Items" ADD FOREIGN KEY ("cart_id") REFERENCES "Carts" ("id");

ALTER TABLE "Cart_Items" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");

ALTER TABLE "Orders" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Order_Lines" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");

ALTER TABLE "Order_Lines" ADD FOREIGN KEY ("order_id") REFERENCES "Orders" ("id");

CREATE INDEX books_array_index ON "Books" USING GIN (tags_array);

Alter table "Reviews" add constraint "user_book_key" unique ("user_id", "book_id");