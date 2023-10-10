CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "first_name" varchar(15),
  "last_name" varchar(15),
  "password" varchar(8),
  "email" varchar(30),
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "machines" (
  "id" integer PRIMARY KEY,
  "name" varchar(15),
  "user_owner" integer,
  "description" varchar,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE "machines" ADD FOREIGN KEY ("user_owner") REFERENCES "users" ("id");
