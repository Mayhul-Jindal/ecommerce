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
  "rating" int not null,
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

-- extension
create extension fuzzystrmatch;

-- mock data
INSERT INTO "Books" ("title", "author", "tags_array", "price", "description", "download_link")
VALUES ('The Great Gatsby', 'F. Scott Fitzgerald', '{1, 2, 3}', 15, 'A classic novel about the American Dream.', 'https://example.com/book1'),
       ('To Kill a Mockingbird', 'Harper Lee', '{4, 5, 6}', 12, 'A powerful exploration of racial injustice.', 'https://example.com/book2'),
       ('1984', 'George Orwell', '{7, 8}', 10, 'A dystopian vision of a totalitarian society.', 'https://example.com/book3'),
       ('Pride and Prejudice', 'Jane Austen', '{9, 10}', 14, 'A timeless tale of romance and social class.', 'https://example.com/book4'),
       ('The Hobbit', 'J.R.R. Tolkien', '{11, 12}', 18, 'An adventure in a fantastical world.', 'https://example.com/book5'),
('Brave New World', 'Aldous Huxley', '{15, 16}', 11, 'A cautionary tale about a controlled society.', 'https://example.com/book7'),
       ('The Lord of the Rings', 'J.R.R. Tolkien', '{11, 12, 17}', 25, 'An epic fantasy trilogy.', 'https://example.com/book8'),
       ('Moby-Dick', 'Herman Melville', '{18, 19}', 16, 'A tale of obsession and revenge.', 'https://example.com/book9'),
       ('Frankenstein', 'Mary Shelley', '{20, 21}', 14, 'A classic horror story.', 'https://example.com/book10'),
       ('The Picture of Dorian Gray', 'Oscar Wilde', '{22, 23}', 17, 'A philosophical exploration of beauty and morality.', 'https://example.com/book11'),
       ('War and Peace', 'Leo Tolstoy', '{24, 25}', 22, 'An epic historical novel.', 'https://example.com/book12'),
       ('The Adventures of Huckleberry Finn', 'Mark Twain', '{26, 27}', 15, 'A journey down the Mississippi River.', 'https://example.com/book13'),
       ('Jane Eyre', 'Charlotte Brontë', '{28, 29}', 12, 'A story of a young woman''s journey to find herself.', 'https://example.com/book14'),
       ('The Odyssey', 'Homer', '{30, 31}', 18, 'An ancient Greek epic.', 'https://example.com/book15'),
       ('Dracula', 'Bram Stoker', '{32, 33}', 16, 'A tale of the legendary vampire.', 'https://example.com/book16'),
       ('The Alchemist', 'Paulo Coelho', '{34, 35}', 10, 'A novel about following your dreams.', 'https://example.com/book17'),
       ('The Scarlet Letter', 'Nathaniel Hawthorne', '{36, 37}', 13, 'A story of sin and redemption.', 'https://example.com/book18'),
       ('Fahrenheit 451', 'Ray Bradbury', '{38, 39}', 11, 'A cautionary tale about censorship.', 'https://example.com/book19'),
       ('The Grapes of Wrath', 'John Steinbeck', '{40, 41}', 14, 'A story of the Great Depression.', 'https://example.com/book20');

INSERT INTO "Tags" ("id", "tag_name") VALUES
  (1, 'Classics'),
  (2, 'Fiction'),
  (3, 'Dystopian'),
  (4, 'Fantasy'),
  (5, 'Romance'),
  (6, 'Young Adult'),
  (7, 'Adventure'),
  (8, 'High Fantasy'),
  (9, 'Science Fiction'),
  (10, 'Literary Fiction'),
  (11, 'Thriller'),
  (12, 'Mystery'),
  (13, 'Coming of Age'),
  (14, 'Horror'),
  (15, 'Magical Realism'),
  (16, 'Philosophical'),
  (17, 'Epic'),
  (18, 'Historical'),
  (19, 'Mystery Thriller'),
  (20, 'Gothic'),
  (21, 'Epic Poetry'),
  (22, 'Historical Fiction'),
  (23, 'Contemporary'),
  (24, 'Paranormal'),
  (25, 'Action'),
  (26, 'Suspense'),
  (27, 'Supernatural'),
  (28, 'Crime'),
  (29, 'Political'),
  (30, 'Comedy'),
  (31, 'Family'),
  (32, 'Tragedy'),
  (33, 'Adventure Fiction'),
  (34, 'Science Fantasy'),
  (35, 'Mythology');