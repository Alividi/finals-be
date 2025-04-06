create table tbl_customer (
	id varchar(50) primary key,
	user_id varchar(50),
	nama_perusahaan varchar(255),
	email_perusahaan varchar(255),
	no_telp_perusahaan varchar(255),
	no_npwp_perusahaan varchar(255),
	provinsi varchar(50),
	kabupaten varchar(50),
	kecamatan varchar(50),
	kelurahan varchar(50),
	alamat varchar(50),
	latitude DOUBLE PRECISION,
	longitude DOUBLE PRECISION
);

create index customer_user_id_fk on tbl_customer(user_id);