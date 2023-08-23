CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "is_admin" bool,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Books" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "title" varchar NOT NULL,
    "author" varchar NOT NULL,
    "genre" varchar NOT NULL,
    "price" decimal NOT NULL,
    "description" text,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Cart" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "book_id" int NOT NULL,
  "quantity" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Purchases" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "book_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Reviews" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "book_id" int NOT NULL,
  "rating" int NOT NULL,
  "comment" text,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Books" ("title");

CREATE INDEX ON "Books" ("author");

CREATE INDEX ON "Books" ("genre");

CREATE INDEX ON "Cart" ("user_id", "book_id");

CREATE INDEX ON "Purchases" ("user_id", "book_id");

CREATE INDEX ON "Reviews" ("user_id", "book_id");

ALTER TABLE "Cart" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Cart" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");

ALTER TABLE "Purchases" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Purchases" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");

ALTER TABLE "Reviews" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Reviews" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");