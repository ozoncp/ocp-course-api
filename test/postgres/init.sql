create database ocp_course_api;
create user ozon_user with encrypted password 'ozon_course';
grant all privileges on database ocp_course_api to ozon_user;

\connect ocp_course_api

create table courses (
    id bigint primary key,
    classroom_id bigint not null,
    name text not null,
    stream text not null
);
alter table courses owner to ozon_user;

create table lessons (
    id bigint primary key,
    course_id bigint not null,
    number integer,
    name text not null
);
alter table lessons owner to ozon_user;
