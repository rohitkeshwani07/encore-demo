-- modify "bookings" table
ALTER TABLE "bookings" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE text, ADD COLUMN "event_page_id" text NULL;
-- drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "bookings_id_seq";
-- modify "calenders" table
ALTER TABLE "calenders" ADD COLUMN "owner_id" text NULL, ADD COLUMN "external_permission" text NULL;
-- modify "events" table
ALTER TABLE "events" ADD COLUMN "owner_id" text NULL, ADD COLUMN "attendees" text NULL, ADD COLUMN "guest" text NULL;
-- create "event_pages" table
CREATE TABLE "event_pages" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "slug" text NULL,
  "title" text NULL,
  "description" text NULL,
  "duration_in_sec" bigint NULL,
  "befor_buffer_time" bigint NULL,
  "after_buffer_time" bigint NULL,
  "expiry" timestamptz NULL,
  "event_type" text NULL,
  "schedule" text NULL,
  "host_id" text NULL,
  PRIMARY KEY ("id")
);
-- create "users" table
CREATE TABLE "users" (
  "id" text NOT NULL,
  "email" text NULL,
  PRIMARY KEY ("id")
);
-- drop "availabilities" table
DROP TABLE "availabilities";
