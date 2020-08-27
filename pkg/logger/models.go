package logger

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// Models defined in logger package
var Models = []interface{}{
	(*EventModel)(nil),
}

// EventModel contain event information logged from a nats event
type EventModel struct {
	ID      uint           `gorm:"primary_key;auto_increment" json:"id"`
	Time    time.Time      `gorm:"column:time"`
	Event   string         `gorm:"column:event"`
	Payload postgres.Jsonb `gorm:"column:payload"`
}
