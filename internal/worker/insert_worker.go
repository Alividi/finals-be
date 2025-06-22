package worker

import (
	"context"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
)

type InsertWorker struct {
	db connection.Connection
}

func NewInsertWorker(db connection.Connection) *InsertWorker {
	return &InsertWorker{db: db}
}

func (w *InsertWorker) Run() {
	ctx := context.Background()
	log.Info().Msg("Running insert worker")

	// Fetch all active services
	var serviceIDs []int64
	query := "SELECT id FROM " + constants.TABLE_SERVICE + " WHERE active = 1"
	err := w.db.Select(ctx, &serviceIDs, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch services")
		return
	}

	now := time.Now()
	ts := now.Format("2006-01-02 15:04:05")

	for _, serviceID := range serviceIDs {
		// Simulate data usage
		// Get last data usage
		var lastUsage float64
		err := w.db.Get(ctx, &lastUsage, `
		SELECT data_usage
		FROM `+constants.TABLE_DATA_USAGE+`
		WHERE service_id = $1
		ORDER BY ts DESC
		LIMIT 1`, serviceID)

		// If no data yet, lastUsage defaults to 0
		if err != nil {
			lastUsage = 0
		}

		// Simulate and accumulate usage
		addedUsage := rand.Float64()*5 + 1 // 1MB - 6MB
		totalUsage := lastUsage + addedUsage

		_, err = w.db.Exec(ctx, `
		INSERT INTO `+constants.TABLE_DATA_USAGE+`
		(service_id, ts, data_usage) VALUES ($1, $2, $3)`,
			serviceID, ts, totalUsage)
		if err != nil {
			log.Error().Err(err).Int64("service_id", serviceID).Msg("Failed to insert data usage")
			continue
		}

		isProblem := false
		var gangguanId *int64

		// 20% chance to generate a problem
		generateProblem := rand.Float64() < 0.2
		var dropRate, signalQuality, obstruction, latency float64
		uptime := &ts
		downlink := rand.Float64()*50 + 10 // 10 - 60 Mbps
		uplink := rand.Float64()*20 + 5    // 5 - 25 Mbps

		if generateProblem {
			switch rand.Intn(7) {
			case 0: // Packet Loss
				dropRate = rand.Float64()*30 + 21
				signalQuality = 100 - dropRate + rand.Float64()*5
				gangguan := int64(1)
				gangguanId = &gangguan
				isProblem = true
			case 1: // Signal Drop
				signalQuality = rand.Float64()*20 + 50 // 50 - 70%
				dropRate = 100 - signalQuality + rand.Float64()*5
				gangguan := int64(2)
				gangguanId = &gangguan
				isProblem = true
			case 2: // Obstruction
				obstruction = rand.Float64()*30 + 21
				gangguan := int64(3)
				gangguanId = &gangguan
				isProblem = true
			case 3: // Offline
				uptime = nil
				gangguan := int64(4)
				gangguanId = &gangguan
				isProblem = true
			case 4: // High Latency
				latency = rand.Float64()*100 + 151
				dropRate = latency / 5
				signalQuality = 100 - latency/2
				gangguan := int64(5)
				gangguanId = &gangguan
				isProblem = true
			case 5: // Bandwidth Drop
				downlink = rand.Float64()*10 + 5 // 5 - 15 Mbps
				uplink = rand.Float64()*5 + 2    // 2 - 7 Mbps
				gangguan := int64(6)
				gangguanId = &gangguan
				isProblem = true
			case 6: // Frequent Reconnects
				uptime = nil
				gangguan := int64(7)
				gangguanId = &gangguan
				isProblem = true
			}
		} else {
			dropRate = rand.Float64() * 10
			signalQuality = rand.Float64()*20 + 80
			obstruction = rand.Float64() * 10
			latency = rand.Float64()*50 + 50
		}

		// Insert telemetry
		_, err = w.db.Exec(ctx, `INSERT INTO `+constants.TABLE_TELEMETRY+`
			(service_id, ts, downlink_troughput, uplink_troughput, ping_drop_rate_avg, ping_latency_ms_avg, obstruction_percent_time, uptime, signal_quality)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			serviceID, ts, downlink, uplink, dropRate, latency, obstruction, uptime, signalQuality)
		if err != nil {
			log.Error().Err(err).Int64("service_id", serviceID).Msg("Failed to insert telemetry")
			continue
		}

		// Update is_problem and gangguan_id in service table
		_, err = w.db.Exec(ctx, `UPDATE `+constants.TABLE_SERVICE+` SET is_problem = $1, gangguan_id = $2 WHERE id = $3`, isProblem, gangguanId, serviceID)
		if err != nil {
			log.Error().Err(err).Int64("service_id", serviceID).Msg("Failed to update service problem status")
			continue
		}

		log.Info().Int64("service_id", serviceID).Msg("Inserted data usage and telemetry")
	}

	log.Info().Msg("Insert worker completed")
}
