-- Insert users

-- 2 Customer Users
INSERT INTO tbl_users (id, username, password, nama, email, no_telp, role, refresh_token)
VALUES 
('u-cust-9382a1ef', 'custuser1', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Customer One', 'cust1@example.com', '0811111111', 'customer', ''),
('u-cust-b762c49d', 'custuser2', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Customer Two', 'cust2@example.com', '0822222222', 'customer', '');

-- 1 Admin User
INSERT INTO tbl_users (id, username, password, nama, email, no_telp, role, refresh_token)
VALUES 
('u-admin-3e8bc45f', 'adminuser', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Admin Guy', 'admin@example.com', '0833333333', 'admin', '');

-- 3 Teknisi Users
INSERT INTO tbl_users (id, username, password, nama, email, no_telp, role, refresh_token)
VALUES 
('u-tech-91f742e2', 'techuser1', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Tech One', 'tech1@example.com', '0844444444', 'teknisi', ''),
('u-tech-7ab1cd33', 'techuser2', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Tech Two', 'tech2@example.com', '0855555555', 'teknisi', ''),
('u-tech-faa30410', 'techuser3', '$2a$10$ONdjrwjCT9Tarpgy/BV39O015NUYQA2uqmdbXaQ5fXKM9plbnxSCG', 'Tech Three', 'tech3@example.com', '0866666666', 'teknisi', '');

-- Insert into tbl_customer
INSERT INTO tbl_customer (id, user_id, nama_perusahaan, email_perusahaan, no_telp_perusahaan, no_npwp_perusahaan)
VALUES 
('cust-b7c312ee', 'u-cust-9382a1ef', 'PT Satu', 'pt1@example.com', '0811111111', '123456789012345'),
('cust-2f74a5c9', 'u-cust-b762c49d', 'PT Dua', 'pt2@example.com', '0822222222', '987654321098765');

-- Insert into tbl_alamat
INSERT INTO tbl_alamat (
	id, customer_id, provinsi, kabupaten, kecamatan, kelurahan,
	rt, rw, alamat, latitude, longitude
) VALUES 
('addr-ec8a5a79', 'cust-b7c312ee', 'Jawa Barat', 'Bandung', 'Coblong', 'Dago', '01', '02', 'Jl. Dago No. 1', -6.893, 107.610),
('addr-4e6f09c2', 'cust-2f74a5c9', 'DKI Jakarta', 'Jakarta Selatan', 'Kebayoran Baru', 'Gandaria', '03', '04', 'Jl. Gandaria No. 2', -6.244, 106.799);


-- Insert into tbl_admin
INSERT INTO tbl_admin (id, user_id, nik, npwp, tgl_lahir)
VALUES 
('adm-18fdc2a0', 'u-admin-3e8bc45f', '3273010000000001', '098765432112345', '1990-01-01');

-- Insert into tbl_teknisi
INSERT INTO tbl_teknisi (id, user_id, status, base)
VALUES 
('tech-a3e9dc41', 'u-tech-91f742e2', 'available', 'Bandung'),
('tech-672cf0aa', 'u-tech-7ab1cd33', 'available', 'Jakarta'),
('tech-57a2b319', 'u-tech-faa30410', 'on-duty', 'Surabaya');
