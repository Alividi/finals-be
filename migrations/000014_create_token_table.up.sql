create table tbl_token (
    id integer primary key,
    user_id integer,
    fcm_token text
);

create index token_user_id_fk on tbl_token(user_id);