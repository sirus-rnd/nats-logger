package logger

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	nats "github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

// QueueName on nats queue group
const QueueName = "nats-logger-service"

// New will create new logger service instance
func New(
	db *gorm.DB,
	natsConn *nats.EncodedConn,
	lg *zap.SugaredLogger,
) *Service {
	return &Service{
		DB: db, Nats: natsConn, Logger: lg,
	}
}

// Service is responsible to log all event in nats message bus to database
type Service struct {
	DB     *gorm.DB
	Nats   *nats.EncodedConn
	Logger *zap.SugaredLogger
	sub    *nats.Subscription
}

// Run service
func (s *Service) Run() error {
	sub, err := s.Nats.QueueSubscribe(">", QueueName, s.LogEvent)
	if err != nil {
		return err
	}
	s.sub = sub

	// loop until closed
	for {
		time.Sleep(time.Second)
	}
}

// Stop service
func (s *Service) Stop() {
	if s.sub != nil {
		s.sub.Unsubscribe()
	}
}

// LogEvent will persist event on database
func (s *Service) LogEvent(m *nats.Msg) {
	payload := postgres.Jsonb{RawMessage: json.RawMessage(m.Data)}
	event := &EventModel{
		Time:    time.Now(),
		Event:   m.Subject,
		Payload: payload,
	}
	s.Logger.Info("receive event: ", event.Event)
	s.Logger.Debug("payload -> ", string(m.Data))

	// save event
	err := s.DB.Save(event).Error
	if err != nil {
		s.Logger.Errorf("failed to save event %s -> %v", event.Event, err)
	}
}
