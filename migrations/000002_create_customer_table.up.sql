create table tbl_customer (
	id varchar(50) primary key,
	user_id varchar(50),
	nama_perusahaan varchar(255),
	email_perusahaan varchar(255),
	no_telp_perusahaan varchar(255),
	no_npwp_perusahaan varchar(255)
);

create index customer_user_id_fk on tbl_customer(user_id);