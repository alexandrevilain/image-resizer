create extension if not exists "uuid-ossp";
create schema api;

create role web_anon nologin;
grant web_anon to supinfo;

create table api.images (
  "id" uuid primary key not null default uuid_generate_v4(),
  "desiredWidth" smallint not null,
  "desiredHeight" smallint not null,
  "originalUrl" varchar(255),
  "resultUrl" varchar(255),
  "finishedat" timestamptz default now()
);

grant usage on schema api to web_anon;
grant select on api.images to web_anon;
