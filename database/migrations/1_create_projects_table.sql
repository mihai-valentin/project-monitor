create table if not exists public.projects
(
    id   integer generated by default as identity,
    name varchar not null
        constraint projects_pk
            unique
);
