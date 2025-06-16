package repository

import (
	"context"
	"database/sql"
	"finals-be/app/services/dto"
	"finals-be/app/services/model"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"finals-be/internal/lib/helper"
	"fmt"
)

type IServiceRepository interface {
	GetServices(ctx context.Context, req dto.ServicesRequest) (services []model.Services, err error)
	GetServiceDetail(ctx context.Context, id int64) (service model.Service, err error)
	GetServiceTelemetry(ctx context.Context, serviceId int64, req dto.ServiceTelemetryRequest) ([]model.Telemetry, error)
	ChangeCoordinate(ctx context.Context, req dto.ChangeCoordinateRequest) error
	GetProblemById(ctx context.Context, gangguanId int64) (problem model.Problem, err error)
	GetStepsByProblemId(ctx context.Context, gangguanId int64) (steps []model.Step, err error)
	GetSubstepsByStepId(ctx context.Context, stepId int64) (substeps []model.Substep, err error)
}

type ServiceRepository struct {
	db connection.Connection
}

func NewServiceRepository(db connection.Connection) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (r *ServiceRepository) GetServices(ctx context.Context, req dto.ServicesRequest) (services []model.Services, err error) {
	query := fmt.Sprintf(
		`SELECT 
			s.id, 
			s.nama_service, 
			s.address_line, 
			s.active, 
			COALESCE(du.data_usage, 0) AS data_usage,
			s.activation_date
		FROM %s s
		INNER JOIN %s c ON c.id = s.customer_id
		LEFT JOIN (
			SELECT DISTINCT ON (service_id) service_id, data_usage
			FROM %s
			ORDER BY service_id, ts DESC
		) du ON du.service_id = s.id
		WHERE c.user_id = $1`,
		constants.TABLE_SERVICE, constants.TABLE_CUSTOMER, constants.TABLE_DATA_USAGE)

	args := []interface{}{req.UserId}
	if req.Active != nil {
		query += " AND s.active = $2"
		args = append(args, *req.Active)
	}

	err = r.db.Select(ctx, &services, query, args...)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (r *ServiceRepository) GetServiceDetail(ctx context.Context, id int64) (model.Service, error) {
	query := fmt.Sprintf(`
		SELECT 
			s.id, s.product_id, s.customer_id, s.gangguan_id, s.nama_service, s.address_line, s.locality, s.latitude, s.longitude, 
			s.service_line_number, s.nickname, s.active, s.ip_kit, s.kit_sn, s.ssid, s.activation_date, s.is_problem,
			p.nama_produk AS device,
			c.nama_perusahaan AS customer_name,
			COALESCE(du.data_usage, 0) AS data_usage
		FROM %s s
		JOIN %s p ON s.product_id = p.id
		JOIN %s c ON s.customer_id = c.id
		LEFT JOIN (
			SELECT DISTINCT ON (service_id) service_id, data_usage
			FROM %s
			ORDER BY service_id, ts DESC
		) du ON du.service_id = s.id
		WHERE s.id = $1
	`, constants.TABLE_SERVICE, constants.TABLE_PRODUK, constants.TABLE_CUSTOMER, constants.TABLE_DATA_USAGE)

	var service model.Service
	err := r.db.Get(ctx, &service, query, id)
	return service, err
}

func (r *ServiceRepository) GetServiceTelemetry(ctx context.Context, serviceId int64, req dto.ServiceTelemetryRequest) ([]model.Telemetry, error) {
	var latestTs string
	getLatestTsQuery := fmt.Sprintf(`SELECT MAX(ts) FROM %s WHERE service_id = $1`, constants.TABLE_TELEMETRY)
	err := r.db.Get(ctx, &latestTs, getLatestTsQuery, serviceId)
	if err != nil {
		return nil, err
	}

	baseQuery := fmt.Sprintf(`
		SELECT 
			id, service_id, ts, downlink_troughput, uplink_troughput, ping_drop_rate_avg, 
			ping_latency_ms_avg, obstruction_percent_time, uptime, signal_quality
		FROM %s
		WHERE service_id = $1
	`, constants.TABLE_TELEMETRY)

	args := []interface{}{serviceId}

	if req.Interval != nil {
		baseQuery += ` AND ts >= ($2)::timestamp - ($3 || ' minutes')::interval`
		args = append(args, latestTs, *req.Interval)
	}

	baseQuery += " ORDER BY ts DESC"

	var telemetry []model.Telemetry
	err = r.db.Select(ctx, &telemetry, baseQuery, args...)
	if err != nil {
		return nil, err
	}

	return telemetry, nil
}

func (r *ServiceRepository) ChangeCoordinate(ctx context.Context, req dto.ChangeCoordinateRequest) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET latitude = $1, longitude = $2
		WHERE id = $3
	`, constants.TABLE_SERVICE)

	_, err := r.db.Exec(ctx, query, req.Latitude, req.Longitude, req.ServiceId)
	return err
}

func (r *ServiceRepository) GetProblemById(ctx context.Context, gangguanId int64) (model.Problem, error) {
	query := fmt.Sprintf(`
		SELECT id, nama_gangguan, deskripsi_gangguan 
		FROM %s
		WHERE id = $1
	`, constants.TABLE_GANGGUAN)
	var problem model.Problem

	err := r.db.Get(ctx, &problem, query, gangguanId)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Problem{}, helper.NewErrNotFound("Gangguan not found")
		}
		return model.Problem{}, err
	}

	return problem, nil
}

func (r *ServiceRepository) GetStepsByProblemId(ctx context.Context, gangguanId int64) ([]model.Step, error) {
	query := fmt.Sprintf(`
		SELECT id, gangguan_id, step, step_number 
		FROM %s
		WHERE gangguan_id = $1
		ORDER BY step_number ASC
	`, constants.TABLE_STEP)
	var steps []model.Step

	err := r.db.Select(ctx, &steps, query, gangguanId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("No steps found for this gangguan")
		}
		return nil, err
	}

	return steps, nil
}

func (r *ServiceRepository) GetSubstepsByStepId(ctx context.Context, stepId int64) ([]model.Substep, error) {
	query := fmt.Sprintf(`
		SELECT id, step_id, substep, gambar, deskripsi 
		FROM %s
		WHERE step_id = $1
		ORDER BY id ASC
	`, constants.TABLE_SUBSTEP)
	var substeps []model.Substep

	err := r.db.Select(ctx, &substeps, query, stepId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("No substeps found for this step")
		}
		return nil, err
	}

	return substeps, nil
}
