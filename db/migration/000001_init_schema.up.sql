CREATE TABLE "users" (
  "id" UUID PRIMARY KEY,
  "full_name" varchar(255),
  "dog" int,
  "address" varchar(255),
  "city" varchar(255),
  "post_code" varchar(255),
  "longitude" varchar(255),
  "latitude" varchar(255),
  "contact" varchar(255),
  "created_at" timestamp DEFAULT 'NOW()'
);

CREATE TABLE "dog" (
  "id" UUID PRIMARY KEY,
  "name" varchar,
  "breed" varchar
);

CREATE TABLE "messages" (
  "id" SERIAL PRIMARY KEY,
  "user_id" UUID,
  "message_body" text,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "messages_receipient" (
  "id" SERIAL PRIMARY KEY,
  "message_id" int,
  "receiver_id" UUID,
  "created_at" timestamp DEFAULT 'now()',
  "status" smallint
);

ALTER TABLE "messages" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "messages_receipient" ADD FOREIGN KEY ("message_id") REFERENCES "messages" ("id");

ALTER TABLE "messages_receipient" ADD FOREIGN KEY ("receiver_id") REFERENCES "users" ("id");