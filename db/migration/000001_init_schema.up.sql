CREATE TABLE "users" (
  "id" UUID PRIMARY KEY NOT NULL,
  "full_name" varchar(255) NOT NULL,
  "contact" varchar(255) NOT NULL,
  "dog" int DEFAULT NULL,
  "address" varchar(255) NOT NULL,
  "city" varchar(255) NOT NULL,
  "post_code" varchar(255) NOT NULL,
  "longitude" varchar(255) NOT NULL,
  "latitude" varchar(255) NOT NULL,
  "created_at" timestamp DEFAULT 'NOW()'
);

CREATE TABLE "dog" (
  "id" UUID PRIMARY KEY NOT NULL,
  "name" varchar,
  "breed" varchar
);

CREATE TABLE "messages" (
  "id" SERIAL PRIMARY KEY,
  "conversation_id" SERIAL,
  "sender_id" UUID,
  "recipient_id" UUID,
  "message_body" text NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "conversation" (
  "id" SERIAL PRIMARY KEY,
  "creator_id" UUID,
  "recipient_id" UUID,
  "created_at" timestamp DEFAULT 'now()',
  "status" smallint DEFAULT 0
);

ALTER TABLE "messages" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "conversation" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("id");

ALTER TABLE "conversation" ADD FOREIGN KEY ("recipient_id") REFERENCES "users" ("id");