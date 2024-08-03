create table public.posts
(
    id         uuid,
    user_id    uuid,
    title      varchar(255) not null,
    body       text         not null,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table public.users
(
    id         uuid,
    first_name varchar(255),
    last_name  varchar(255),
    email      varchar(255),
    password   varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

alter table public.users
    owner to postgres;

alter table public.posts
    owner to postgres;

