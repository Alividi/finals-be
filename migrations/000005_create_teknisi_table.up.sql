create table tbl_teknisi (
	id integer primary key,
	user_id integer,
	status varchar(50),
	base varchar(50)
);

create index teknisi_user_id_fk on tbl_teknisi(user_id);