package helper

import (
	"fmt"
	"time"
)

const ticketPrefix = "TKT"

func GenerateTicketNumber(serviceId, gangguanId, customerId int64) string {
	timestamp := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS
	return fmt.Sprintf("%s-S%dG%dC%dT%s", ticketPrefix, serviceId, gangguanId, customerId, timestamp)
}
