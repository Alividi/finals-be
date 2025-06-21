create table tbl_users (
	id integer primary key,
	username varchar(255),
	password varchar(255),
	nama varchar(50),
	email varchar(50),
	no_telp varchar(50),
	role varchar(50),
	refresh_token varchar(510)
);