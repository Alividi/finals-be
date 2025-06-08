-- Insert users

-- 2 Customer Users
INSERT INTO tbl_users (id, username, password, nama, email, no_telp, role, refresh_token)
VALUES 
(1, 'custuser1', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Customer One', 'cust1@example.com', '0811111111', 'customer', ''),
(2, 'custuser2', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Customer Two', 'cust2@example.com', '0822222222', 'customer', '');

-- 1 Admin User
INSERT INTO tbl_users (id, username, password, nama, email, no_telp, role, refresh_token)
VALUES 
(3, 'adminuser', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Admin Guy', 'admin@example.com', '0833333333', 'admin', '');

-- 3 Teknisi Users
INSERT INTO tbl_users (id, username, password, nama, email, no_telp, role, refresh_token)
VALUES 
(4, 'techuser1', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Tech One', 'tech1@example.com', '0844444444', 'teknisi', ''),
(5, 'techuser2', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Tech Two', 'tech2@example.com', '0855555555', 'teknisi', ''),
(6, 'techuser3', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Tech Three', 'tech3@example.com', '0866666666', 'teknisi', '');

-- Insert into tbl_customer
INSERT INTO tbl_customer (id, user_id, nama_perusahaan, email_perusahaan, no_telp_perusahaan, no_npwp_perusahaan)
VALUES 
(1, 1, 'PT Satu', 'pt1@example.com', '0811111111', '123456789012345'),
(2, 2, 'PT Dua', 'pt2@example.com', '0822222222', '987654321098765');

-- Insert into tbl_alamat
INSERT INTO tbl_alamat (
	id, customer_id, provinsi, kabupaten, kecamatan, kelurahan,
	rt, rw, alamat, latitude, longitude
) VALUES 
(1, 1, 'Jawa Barat', 'Bandung', 'Coblong', 'Dago', '01', '02', 'Jl. Dago No. 1', -6.893, 107.610),
(2, 2, 'DKI Jakarta', 'Jakarta Selatan', 'Kebayoran Baru', 'Gandaria', '03', '04', 'Jl. Gandaria No. 2', -6.244, 106.799);


-- Insert into tbl_admin
INSERT INTO tbl_admin (id, user_id, nik, npwp, tgl_lahir)
VALUES 
(1, 3, '3273010000000001', '098765432112345', '1990-01-01');

-- Insert into tbl_teknisi
INSERT INTO tbl_teknisi (id, user_id, status, base)
VALUES 
(1, 4, 'available', 'Bandung'),
(2, 5, 'available', 'Jakarta'),
(3, 6, 'on-duty', 'Surabaya');
