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

alter table public.posts
    owner to postgres;

