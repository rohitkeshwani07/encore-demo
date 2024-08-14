-- reverse: modify "event_pages" table
ALTER TABLE "event_pages" ADD COLUMN "after_buffer_time" bigint NULL, ADD COLUMN "befor_buffer_time" bigint NULL;
