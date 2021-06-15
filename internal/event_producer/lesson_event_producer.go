package event_producer

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/cenkalti/backoff/v4"
	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
)

type LessonEventProducer interface {
	SendEvent(*model.LessonEvent) error
	Close() error
}

type lessonEventProducer struct {
	topic    string
	producer sarama.SyncProducer
}

func (p *lessonEventProducer) prepareMessage(evt *model.LessonEvent) (*sarama.ProducerMessage, error) {
	encoded, err := json.Marshal(evt)
	if err != nil {
		return nil, err
	}
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(string(encoded)),
	}
	log.Debug().Msgf("event: %v", evt)
	log.Debug().Msgf("message: %v", msg)
	return msg, nil
}

func (p *lessonEventProducer) SendEvent(evt *model.LessonEvent) error {
	msg, err := p.prepareMessage(evt)
	if err != nil {
		return err
	}
	_, _, err = p.producer.SendMessage(msg)
	return err
}

func (p *lessonEventProducer) Close() error {
	return p.producer.Close()
}

func NewLessonEventProducer(
	topic string,
	brokers []string,
) (*lessonEventProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 10
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &lessonEventProducer{topic: topic, producer: producer}, nil
}

func MakeLessonEventProducerWithRetry(
	ctx context.Context,
	topic string,
	brokers []string,
) (*lessonEventProducer, error) {
	var p *lessonEventProducer
	err := backoff.Retry(func() error {
		var err error
		p, err = NewLessonEventProducer(topic, brokers)
		if err != nil {
			log.Debug().Err(err).Msg("Attempt to connect to Kafka failed")
			return err
		}
		return err
	},
		backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		return nil, err
	}
	return p, nil
}
