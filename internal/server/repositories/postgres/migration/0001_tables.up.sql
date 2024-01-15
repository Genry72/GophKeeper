/*
   users
 */
create table if not exists users
(
    id       bigserial primary key   not null,
    username      varchar                 not null unique,
    password_hash varchar                 not null,
    created_at    timestamp default now() not null
);

comment on table users is 'Пользователи системы';
comment on column users.id is 'Уникальный идентификатор пользователя';
comment on column users.username is 'Логин пользователя';
comment on column users.password_hash is 'Пароль пользователя';
comment on column users.created_at is 'Дата создания';