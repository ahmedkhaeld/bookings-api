CREATE TABLE "users" (
                            "id" integer PRIMARY KEY,
                            "first_name" varchar NOT NULL DEFAULT '',
                            "last_name" varchar NOT NULL DEFAULT '',
                            "email" varchar NOT NULL,
                            "password" varchar(60) NOT NULL ,
                            "access_level" varchar NOT NULL DEFAULT 1
);


CREATE TABLE "rooms" (
                     id integer primary key ,
                     room_name varchar default ''
);


CREATE TABLE "reservations" (
    "id" integer PRIMARY KEY,
    "first_name" varchar NOT NULL default '',
    "last_name" varchar not null default '',
    "email" varchar not null ,
    "phone" varchar not null default '',
    "start_date" timestamp not null default (now()),
    "end_date" timestamp not null default (now()),
    "processed" integer default 0,
    "room_id" integer references rooms(id)
);

CREATE TABLE "restrictions" (
    "id" integer primary key,
    restriction_name varchar default ''
);

CREATE TABLE "room_restrictions" (
    "id" integer primary key,
    start_date timestamp not null default (now()),
    end_date timestamp not null default (now()),
    "room_id" integer references rooms(id),
    reservation_id integer not null references reservations(id) on update cascade on delete cascade ,
    restriction_id integer references restrictions(id)

);

CREATE unique INDEX users_emails_idx ON "users" ("email");
create index reservations_emails_idx on reservations(email);
create index reservations_last_name_idx on reservations(last_name);
create index reservations_id_idx on room_restrictions(reservation_id);
create index room_id_idx on room_restrictions(room_id);
create index start_end_idx on room_restrictions(start_date,end_date);
