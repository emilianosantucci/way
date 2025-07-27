-- create schemas
create schema way;
alter schema way owner to way;

-- add support for uuid
create extension if not exists "uuid-ossp" schema way;