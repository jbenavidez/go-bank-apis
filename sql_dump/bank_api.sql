-- -------------------------------------------------------------
-- TablePlus 5.6.0(514)
--
-- https://tableplus.com/
--
-- Database: bank
-- Generation Time: 2025-11-16 14:29:11.3420
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."accounts";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS accounts_id_seq;

-- Table Definition
CREATE TABLE "public"."accounts" (
    "id" int4 NOT NULL DEFAULT nextval('accounts_id_seq'::regclass),
    "user_id" int4 NOT NULL,
    "acc_type" varchar(255),
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "amount" numeric(10,2) NOT NULL,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "first_name" varchar(255),
    "last_name" varchar(255),
    "email" varchar(255),
    "username" varchar(255),
    PRIMARY KEY ("id")
);

INSERT INTO "public"."accounts" ("id", "user_id", "acc_type", "created_at", "updated_at", "amount") VALUES
(1, 1, 'checking', '2025-11-15 19:23:34.834804+00', '2025-11-16 19:25:30.562509+00', 8802.00),
(2, 1, 'checking', '2025-11-16 19:13:10.573853+00', '2025-11-16 19:13:10.573854+00', 5001.00);

INSERT INTO "public"."users" ("id", "first_name", "last_name", "email", "username") VALUES
(1, 'Janet', 'Doe', 'janet@doe.com', 'valinor_elf_hello'),
(2, 'John', 'Doe', 'johndoe@doe.com', 'Gondor_elf'),
(3, 'Joh1n', 'Doe', 'johndoe@doe.com', 'Gondor_elf');

ALTER TABLE "public"."accounts" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE ON UPDATE CASCADE;
