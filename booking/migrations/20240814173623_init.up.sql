-- create "availabilities" table
CREATE TABLE "availabilities" (
  "start" text NULL,
  "end" text NULL
);
-- create "bookings" table
CREATE TABLE "bookings" (
  "id" bigserial NOT NULL,
  "start" timestamptz NULL,
  "end" timestamptz NULL,
  "email" text NULL,
  PRIMARY KEY ("id")
);
