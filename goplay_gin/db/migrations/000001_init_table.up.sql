CREATE TABLE "author" (
  "id" varchar PRIMARY KEY,
  "firstname" varchar(100) NOT NULL,
  "secondname" varchar(100) NOT NULL
);

CREATE TABLE "book" (
  "id" varchar PRIMARY KEY,
  "author_id" varchar NOT NULL,
  "isbn" varchar(200) NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE "book" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");