package model

type Service struct {
	ID                int64   `db:"id"`
	FkProductId       int64   `db:"product_id"`
	FkCustomerId      int64   `db:"customer_id"`
	CustomerName      string  `json:"customer_name" db:"customer_name"`
	FkGangguanId      *int64  `db:"gangguan_id"`
	NamaService       string  `db:"nama_service"`
	AddressLine       string  `db:"address_line"`
	Locality          string  `db:"locality"`
	Latitude          float64 `db:"latitude"`
	Longitude         float64 `db:"longitude"`
	ServiceLineNumber string  `db:"service_line_number"`
	Nickname          string  `db:"nickname"`
	Active            int64   `db:"active"`
	Ipkit             int64   `db:"ip_kit"`
	KitSn             string  `db:"kit_sn"`
	SSID              string  `db:"ssid"`
	ActivationDate    string  `db:"activation_date"`
	IsProblem         bool    `db:"is_problem"`
	Device            string  `json:"device" db:"device"`
	DataUsage         float64 `json:"data_usage" db:"data_usage"`
}

type Services struct {
	ID             int64   `json:"id" db:"id"`
	NamaService    string  `json:"nama_service" db:"nama_service"`
	AddressLine    string  `json:"address_line" db:"address_line"`
	Active         int64   `json:"active" db:"active"`
	DataUsage      float64 `json:"data_usage" db:"data_usage"`
	ActivationDate string  `json:"activation_date" db:"activation_date"`
}

type Telemetry struct {
	ID                int64   `db:"id"`
	FkServiceId       int64   `db:"service_id"`
	Timestamp         string  `db:"ts"`
	DownlinkTroughput float64 `db:"downlink_troughput"`
	UplinkTroughput   float64 `db:"uplink_troughput"`
	PingDropRate      float64 `db:"ping_drop_rate_avg"`
	Latency           float64 `db:"ping_latency_ms_avg"`
	Obstruction       float64 `db:"obstruction_percent_time"`
	Uptime            string  `db:"uptime"`
	SignalQuality     float64 `db:"signal_quality"`
}

type DataUsage struct {
	ID          int64   `db:"id"`
	FkServiceId int64   `db:"service_id"`
	Timestamp   string  `db:"ts"`
	DataUsage   float64 `db:"data_usage"`
}
