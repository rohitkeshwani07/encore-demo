-- create "calenders" table
CREATE TABLE "calenders" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "external_id" text NULL,
  "external_type" text NULL,
  PRIMARY KEY ("id")
);
-- create "events" table
CREATE TABLE "events" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "calender_id" text NULL,
  "start_time" timestamptz NULL,
  "end_time" timestamptz NULL,
  "seq" smallint NULL,
  "event_payload" text NULL,
  "i_cal_payload" text NULL,
  "i_cal_uid" text NULL,
  "external_id" text NULL,
  "external_type" text NULL,
  PRIMARY KEY ("id")
);
