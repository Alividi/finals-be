-- Insert initial data into tbl_service
INSERT INTO tbl_service (
    id, product_id, customer_id, nama_service, address_line, locality,
    latitude, longitude, service_line_number, nickname, active,
    ip_kit, kit_sn, ssid, activation_date, is_problem
) VALUES
(1, 1, 1, 'Service Fixed 50GB', 'Jl. Merdeka No. 1', 'Jakarta', -6.2088, 106.8456, 'SLN-001', 'Site Jakarta', 1, '192.168.1', 'KIT-001', 'SSID-JKT', '2025-06-01 10:01:00', FALSE),
(2, 1, 2, 'Service Fixed 1TB', 'Jl. Malioboro No. 99', 'Yogyakarta', -7.8014, 110.3644, 'SLN-002', 'Site Yogyakarta', 1, '192.168.2', 'KIT-002', 'SSID-YOG', '2025-06-02 11:31:00', FALSE),
(3, 2, 1, 'Service Land 50GB', 'Jl. Asia Afrika No. 88', 'Bandung', -6.9214, 107.6079, 'SLN-003', 'Site Bandung', 1, '192.168.3', 'KIT-003', 'SSID-BDG', '2025-06-03 12:01:00', FALSE),
(4, 2, 2, 'Service Land 1TB', 'Jl. Sudirman No. 45', 'Jakarta', -6.2088, 106.8456, 'SLN-004', 'Site Jakarta Sudirman', 1, '192.168.4', 'KIT-004', 'SSID-JKT-SUD', '2025-06-04 13:01:00', FALSE);

-- Insert initial data into tbl_datausage
-- Data Usage for Service ID 1 (SLN-001)
INSERT INTO tbl_datausage ( service_id, data_usage, ts) VALUES
( 1, 0.5, '2025-06-01 10:01:00'),
( 1, 1.0, '2025-06-01 10:02:00'),
( 1, 1.6, '2025-06-01 10:03:00'),
( 1, 2.3, '2025-06-01 10:04:00'),
( 1, 3.1, '2025-06-01 10:05:00'),
( 1, 4.0, '2025-06-01 10:06:00'),
( 1, 5.0, '2025-06-01 10:07:00');

-- Data Usage for Service ID 2 (SLN-002)
INSERT INTO tbl_datausage ( service_id, data_usage, ts) VALUES
( 2, 0.4, '2025-06-02 11:31:00'),
( 2, 0.9, '2025-06-02 11:32:00'),
( 2, 1.5, '2025-06-02 11:33:00'),
( 2, 2.2, '2025-06-02 11:34:00'),
( 2, 3.0, '2025-06-02 11:35:00'),
( 2, 3.9, '2025-06-02 11:36:00'),
( 2, 4.8, '2025-06-02 11:37:00');

-- Data Usage for Service ID 3 (SLN-003)
INSERT INTO tbl_datausage ( service_id, data_usage, ts) VALUES
( 3, 0.3, '2025-06-03 12:01:00'),
( 3, 0.8, '2025-06-03 12:02:00'),
( 3, 1.4, '2025-06-03 12:03:00'),
( 3, 2.1, '2025-06-03 12:04:00'),
( 3, 2.9, '2025-06-03 12:05:00'),
( 3, 3.8, '2025-06-03 12:06:00'),
( 3, 4.7, '2025-06-03 12:07:00');

-- Data Usage for Service ID 4 (SLN-004)
INSERT INTO tbl_datausage ( service_id, data_usage, ts) VALUES
( 4, 0.6, '2025-06-04 13:01:00'),
( 4, 1.2, '2025-06-04 13:02:00'),
( 4, 1.8, '2025-06-04 13:03:00'),
( 4, 2.5, '2025-06-04 13:04:00'),
( 4, 3.3, '2025-06-04 13:05:00'),
( 4, 4.1, '2025-06-04 13:06:00'),
( 4, 4.9, '2025-06-04 13:07:00');

-- Insert initial data into tbl_telemetry
-- Telemetry for Service ID 1 (SLN-001)
INSERT INTO tbl_telemetry (
     service_id, ts, downlink_troughput, uplink_troughput,
    ping_drop_rate_avg, ping_latency_ms_avg, obstruction_percent_time, uptime, signal_quality
) VALUES
( 1, '2025-06-01 10:01:00', 45.0, 11.5, 0.01, 30.2, 4.5, '2025-06-01 09:01:00', 97.0),
( 1, '2025-06-01 10:02:00', 44.8, 11.6, 0.01, 30.5, 4.4, '2025-06-01 09:01:00', 97.2),
( 1, '2025-06-01 10:03:00', 44.6, 11.7, 0.02, 30.3, 4.6, '2025-06-01 09:01:00', 96.9),
( 1, '2025-06-01 10:04:00', 44.2, 11.8, 0.01, 30.1, 4.5, '2025-06-01 09:01:00', 96.8),
( 1, '2025-06-01 10:05:00', 43.9, 11.4, 0.01, 29.9, 4.3, '2025-06-01 09:01:00', 96.7);

-- Telemetry for Service ID 2 (SLN-002)
INSERT INTO tbl_telemetry (
     service_id, ts, downlink_troughput, uplink_troughput,
    ping_drop_rate_avg, ping_latency_ms_avg, obstruction_percent_time, uptime, signal_quality
) VALUES
( 2, '2025-06-02 11:31:00', 30.0, 7.2, 0.03, 45.0, 10.0, '2025-06-02 10:31:00', 85.0),
( 2, '2025-06-02 11:32:00', 29.8, 7.1, 0.03, 45.2, 10.2, '2025-06-02 10:31:00', 84.8),
( 2, '2025-06-02 11:33:00', 29.5, 7.0, 0.04, 45.4, 10.5, '2025-06-02 10:31:00', 84.6),
( 2, '2025-06-02 11:34:00', 29.3, 6.9, 0.04, 45.6, 10.6, '2025-06-02 10:31:00', 84.3),
( 2, '2025-06-02 11:35:00', 29.0, 6.8, 0.05, 45.8, 10.8, '2025-06-02 10:31:00', 84.0);

-- Telemetry for Service ID 3 (SLN-003)
INSERT INTO tbl_telemetry (
     service_id, ts, downlink_troughput, uplink_troughput,
    ping_drop_rate_avg, ping_latency_ms_avg, obstruction_percent_time, uptime, signal_quality
) VALUES
( 3, '2025-06-03 12:01:00', 20.0, 5.0, 0.02, 60.0, 15.0, '2025-06-03 11:01:00', 75.0),
( 3, '2025-06-03 12:02:00', 19.8, 4.9, 0.02, 60.2, 15.2, '2025-06-03 11:01:00', 74.8),
( 3, '2025-06-03 12:03:00', 19.5, 4.8, 0.03, 60.4, 15.5, '2025-06-03 11:01:00', 74.6),
( 3, '2025-06-03 12:04:00', 19.3, 4.7, 0.03, 60.6, 15.6, '2025-06-03 11:01:00', 74.3),
( 3, '2025-06-03 12:05:00', 19.0, 4.6, 0.04, 60.8, 15.8, '2025-06-03 11:01:00', 74.0);

-- Telemetry for Service ID 4 (SLN-004)
INSERT INTO tbl_telemetry (
     service_id, ts, downlink_troughput, uplink_troughput,
    ping_drop_rate_avg, ping_latency_ms_avg, obstruction_percent_time, uptime, signal_quality
) VALUES
( 4, '2025-06-04 13:01:00', 25.0, 6.0, 0.01, 50.0, 12.0, '2025-06-04 12:01:00', 80.0),
( 4, '2025-06-04 13:02:00', 24.8, 5.9, 0.01, 50.2, 12.2, '2025-06-04 12:01:00', 79.8),
( 4, '2025-06-04 13:03:00', 24.5, 5.8, 0.02, 50.4, 12.5, '2025-06-04 12:01:00', 79.6),
( 4, '2025-06-04 13:04:00', 24.3, 5.7, 0.02, 50.6, 12.6, '2025-06-04 12:01:00', 79.3),
( 4, '2025-06-04 13:05:00', 24.0, 5.6, 0.03, 50.8, 12.8, '2025-06-04 12:01:00', 79.0);
