-- TODO: what is the use case if this in production

CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "is_admin" boolean not null,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Books" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "title" varchar NOT NULL,
    "author" varchar NOT NULL,
    "tags" text[] NOT NULL,
    "price" int NOT NULL,
    "quantity" int not null,
    "description" varchar not null,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Cart" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "book_id" bigint NOT NULL,
  "quantity" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Purchases" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "book_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Reviews" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "book_id" bigint NOT NULL,
  "rating" int NOT NULL,
  "comment" varchar not null,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Books" ("title");

CREATE INDEX ON "Books" ("author");

CREATE INDEX books_x ON "Books" USING GIN (tags);

CREATE INDEX ON "Cart" ("user_id", "book_id");

CREATE INDEX ON "Purchases" ("user_id", "book_id");

CREATE INDEX ON "Reviews" ("user_id", "book_id");

ALTER TABLE "Cart" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Cart" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");

ALTER TABLE "Purchases" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Purchases" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");

ALTER TABLE "Reviews" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

ALTER TABLE "Reviews" ADD FOREIGN KEY ("book_id") REFERENCES "Books" ("id");