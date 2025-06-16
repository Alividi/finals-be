package dto

type GetServicesResponse struct {
	ID             int64   `json:"id" db:"id"`
	NamaService    string  `json:"nama_service" db:"nama_service"`
	AddressLine    string  `json:"address_line" db:"address_line"`
	Active         int64   `json:"active" db:"active"`
	DataUsage      float64 `json:"data_usage" db:"data_usage"`
	ActivationDate string  `json:"activation_date" db:"activation_date"`
}

type ServicesRequest struct {
	UserId     int64  `json:"user_id" validate:"required"`
	CustomerID int64  `json:"-"`
	Active     *int64 `json:"active" validate:"omitempty"`
}

type GetServiceDetailResponse struct {
	ID                int64   `json:"id" db:"id"`
	ProductId         int64   `json:"product_id" db:"product_id"`
	CustomerId        int64   `json:"customer_id" db:"customer_id"`
	CustomerName      string  `json:"customer_name" db:"nama_perusahaan"`
	GangguanId        *int64  `json:"gangguan_id" db:"gangguan_id"`
	NamaService       string  `json:"nama_service" db:"nama_service"`
	AddressLine       string  `json:"address_line" db:"address_line"`
	Locality          string  `json:"locality" db:"locality"`
	Latitude          float64 `json:"latitude" db:"latitude"`
	Longitude         float64 `json:"longitude" db:"longitude"`
	ServiceLineNumber string  `json:"service_line_number" db:"service_line_number"`
	Nickname          string  `json:"nickname" db:"nickname"`
	Active            int64   `json:"active" db:"active"`
	Ipkit             int64   `json:"ip_kit" db:"ip_kit"`
	KitSn             string  `json:"kit_sn" db:"kit_sn"`
	SSID              string  `json:"ssid" db:"ssid"`
	ActivationDate    string  `json:"activation_date" db:"activation_date"`
	IsProblem         bool    `json:"is_problem" db:"is_problem"`
	Device            string  `json:"device" db:"nama_produk"`
	DataUsage         float64 `json:"data_usage" db:"data_usage"`
}

type GetServiceTelemetryResponse struct {
	ID                int64   `json:"id" db:"id"`
	ServiceId         int64   `json:"service_id" db:"service_id"`
	Timestamp         string  `json:"ts" db:"ts"`
	DownlinkTroughput float64 `json:"downlink_troughput" db:"downlink_troughput"`
	UplinkTroughput   float64 `json:"uplink_troughput" db:"uplink_troughput"`
	PingDropRate      float64 `json:"ping_drop_rate_avg" db:"ping_drop_rate_avg"`
	Latency           float64 `json:"ping_latency_ms_avg" db:"ping_latency_ms_avg"`
	Obstruction       float64 `json:"obstruction_percent_time" db:"obstruction_percent_time"`
	Uptime            string  `json:"uptime" db:"uptime"`
	SignalQuality     float64 `json:"signal_quality" db:"signal_quality"`
}

type ServiceTelemetryRequest struct {
	Interval *int64 `json:"interval" validate:"omitempty"`
}

type ChangeCoordinateRequest struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	ServiceId int64   `json:"service_id" validate:"required"`
}

type Substep struct {
	SubstepId int64  `json:"substep_id" db:"substep_id"`
	Substep   string `json:"substep" db:"substep"`
	Gambar    string `json:"gambar" db:"gambar"`
	Deskripsi string `json:"deskripsi" db:"deskripsi"`
}

type Step struct {
	StepId     int64     `json:"step_id" db:"step_id"`
	Step       string    `json:"step" db:"step"`
	StepNumber int64     `json:"step_number" db:"step_number"`
	Substeps   []Substep `json:"substeps" db:"substeps"`
}

type GetSolutionResponse struct {
	GangguanId        int64  `json:"gangguan_id" db:"gangguan_id"`
	NamaGangguan      string `json:"nama_gangguan" db:"nama_gangguan"`
	DeskripsiGangguan string `json:"deskripsi_gangguan" db:"deskripsi_gangguan"`
	Steps             []Step `json:"steps" db:"steps"`
}
