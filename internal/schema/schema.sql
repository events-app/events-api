BEGIN;
--
-- Create model Card
--
CREATE TABLE "card" ("id" serial NOT NULL PRIMARY KEY, "name" varchar(200) NOT NULL, "text" text NOT NULL, "created" timestamp NOT NULL, "updated" timestamp NOT NULL);
--
-- Create model File
--
CREATE TABLE "file" ("id" serial NOT NULL PRIMARY KEY, "name" varchar(200) NOT NULL, "path" varchar(500) NOT NULL, "created" timestamp NOT NULL, "updated" timestamp NOT NULL);
--
-- Create model User
--
CREATE TABLE "user" ("id" serial NOT NULL PRIMARY KEY, "username" varchar(200) NOT NULL, "password" varchar(100) NOT NULL, "role" varchar(100) NOT NULL, "created" timestamp NOT NULL, "updated" timestamp NOT NULL);
--
-- Create model Menu
--
CREATE TABLE "menu" ("id" serial NOT NULL PRIMARY KEY, "name" varchar(200) NOT NULL, "created" timestamp NOT NULL, "updated" timestamp NOT NULL, "card_id" integer NOT NULL);
ALTER TABLE "menu" ADD CONSTRAINT "menu_card_id_7cfd48a2_fk_card_id" FOREIGN KEY ("card_id") REFERENCES "card" ("id") DEFERRABLE INITIALLY DEFERRED;
CREATE INDEX "menu_card_id_7cfd48a2" ON "menu" ("card_id");
COMMIT;
