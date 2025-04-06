create table tbl_alamat (
	id varchar(50) primary key,
	customer_id varchar(50),
	provinsi varchar(50),
	kabupaten varchar(50),
	kecamatan varchar(50),
	kelurahan varchar(50),
	rt varchar(50),
	rw varchar(50),
	alamat varchar(50),
	latitude DOUBLE PRECISION,
	longitude DOUBLE PRECISION
);

create index customer_id_fk on tbl_alamat(customer_id);