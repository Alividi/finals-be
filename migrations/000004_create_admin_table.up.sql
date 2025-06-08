create table tbl_admin (
	id integer primary key,
	user_id integer,
	nik varchar(50),
	npwp varchar(50),
	tgl_lahir timestamp
);

create index admin_user_id_fk on tbl_admin(user_id);