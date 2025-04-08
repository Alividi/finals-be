create table tbl_admin (
	id varchar(50) primary key,
	user_id varchar(50),
	nik varchar(50),
	npwp varchar(50),
	tgl_lahir timestamp
);

create index admin_user_id_fk on tbl_admin(user_id);