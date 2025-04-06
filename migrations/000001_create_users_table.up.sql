create table tbl_users (
	id varchar(50) primary key,
	username varchar(255),
	password varchar(255),
	role varchar(50),
	refresh_token varchar(255)
);