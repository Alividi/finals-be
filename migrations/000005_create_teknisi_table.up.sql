create table tbl_teknisi (
	id varchar(50) primary key,
	user_id varchar(50),
	nama varchar(50),
	email varchar(50),
	no_telp varchar(50),
	status varchar(50),
	base varchar(50)
);

create index teknisi_user_id_fk on tbl_teknisi(user_id);