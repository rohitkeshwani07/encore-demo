-- reverse: drop "availabilities" table
CREATE TABLE "availabilities" (
  "start" text NULL,
  "end" text NULL
);
-- reverse: create "users" table
DROP TABLE "users";
-- reverse: create "event_pages" table
DROP TABLE "event_pages";
-- reverse: modify "events" table
ALTER TABLE "events" DROP COLUMN "guest", DROP COLUMN "attendees", DROP COLUMN "owner_id";
-- reverse: modify "calenders" table
ALTER TABLE "calenders" DROP COLUMN "external_permission", DROP COLUMN "owner_id";
-- reverse: drop sequence used by serial column "id"
CREATE SEQUENCE IF NOT EXISTS "bookings_id_seq" OWNED BY "bookings"."id";
-- reverse: modify "bookings" table
ALTER TABLE "bookings" DROP COLUMN "event_page_id", ALTER COLUMN "id" SET DEFAULT nextval('"bookings_id_seq"'), ALTER COLUMN "id" TYPE bigint;
