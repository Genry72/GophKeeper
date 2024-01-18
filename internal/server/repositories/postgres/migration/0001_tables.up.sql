/*
   Таблица пользователей
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

/*
   Типы хранимых секретов
 */
create table secret_types
(
    id  bigserial primary key   not null,
    name text    not null
);

comment on table secret_types is 'Типы хранимой информации';
comment on column secret_types.id is 'Уникальный идентификатор типа хранимой информации';
comment on column secret_types.name is 'Название типа';

insert into secret_types (name) values ('login/password');
insert into secret_types (name) values ('text');
insert into secret_types (name) values ('binary');
insert into secret_types (name) values ('bank_card');

/*
   Секреты пользователей
 */

create table user_secrets
(
    id             bigserial primary key   not null,
    user_id        bigint                  not null
        constraint user_secrets_users_id_fk
            references users,
    secret_type_id bigint                  not null
        constraint user_secrets_secret_types_id_fk
            references secret_types,
    name           text                    not null,
    data           integer,
    created_at     timestamp default now() not null,
    updated_at     timestamp default now() not null,
    deleted_at     timestamp
);

comment on table user_secrets is 'Секреты пользователей';
comment on column user_secrets.id is 'Уникальный идентификатор секрета';
comment on column user_secrets.user_id is 'ID владельца секрета из таблицы users';
comment on column user_secrets.secret_type_id is 'ID типа секрета из таблицы secret_types';
comment on column user_secrets.name is 'Имя секрета';
comment on column user_secrets.data is 'Закодированные данные с секретом';
comment on column user_secrets.created_at is 'Дата создания секрета';
comment on column user_secrets.updated_at is 'Дата изменения секрета';
comment on column user_secrets.deleted_at is 'Дата удаления секрета';