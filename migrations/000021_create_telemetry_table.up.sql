CREATE TABLE tbl_telemetry (
    id integer PRIMARY KEY,
    service_line_number VARCHAR(50),
    service_id integer,
    ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    downlink_troughput float,
    uplink_troughput float,
    ping_drop_rate_avg float,
    ping_latency_ms_avg float,
    obstruction_percent_time float,
    uptime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    signal_quality float,
    FOREIGN KEY (service_line_number) REFERENCES tbl_service_line(service_line_number),
    FOREIGN KEY (service_id) REFERENCES tbl_service(id)
);

CREATE INDEX telemetry_service_line_number_fk ON tbl_telemetry(service_line_number);
CREATE INDEX telemetry_service_id_fk ON tbl_telemetry(service_id);