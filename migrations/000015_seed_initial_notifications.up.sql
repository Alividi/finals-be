-- Insert into tbl_notifikasi
INSERT INTO tbl_notifikasi (user_id, is_read, judul, type, deskripsi)
VALUES 
-- Notifications for Customer Users
( 1, false, 'Welcome', 'info', 'Welcome to our platform, Customer One!'),
( 2, true, 'Profile Update', 'update', 'Your profile has been successfully updated.'),

-- Notifications for Admin User
( 3, false, 'System Notice', 'alert', 'System maintenance will occur this weekend.'),
( 3, false, 'New User Registration', 'info', 'A new user has registered on the platform.'),

-- Notifications for Teknisi Users
( 4, false, 'New Assignment', 'task', 'You have been assigned a new maintenance task in Bandung.'),
( 5, true, 'Reminder', 'reminder', 'Please submit your weekly report.'),
( 6, false, 'Task Completed', 'success', 'The system has recorded your last completed task.');
