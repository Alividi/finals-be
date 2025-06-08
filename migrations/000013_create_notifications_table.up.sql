DROP TABLE IF EXISTS tbl_notifikasi;

create table tbl_notifikasi (
    id integer primary key,
    user_id integer,
    is_read boolean default false,
    judul varchar(255),
    type varchar(50),
    deskripsi text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    foreign key (user_id) references tbl_users(id) on delete cascade
);

create index notfikasi_users_id_fk on tbl_notifikasi(user_id);