CREATE TABLE tbl_token (
    id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id integer,
    fcm_token text,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    foreign key (user_id) references tbl_users(id) on delete cascade
);

CREATE INDEX token_users_id_fk ON tbl_token(user_id);