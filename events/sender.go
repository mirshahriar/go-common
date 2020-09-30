package events

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"time"
)

const (
	producerRetryPeriod = 5 * time.Second
	producerMaxRetries = 5
)

type EventProducer interface {
	Send(ctx context.Context, event Event) error
}

var _ EventProducer = &KafkaCloudEventsProducer{}
type KafkaCloudEventsProducer struct {
	client cloudevents.Client
	source string
}

func NewKafkaCloudEventsProducer(config *CloudEventsConfig) (*KafkaCloudEventsProducer, error) {
	sender, err := kafka_sarama.NewSender(config.KafkaBrokers, config.SaramaConfig, config.GetPrefixedTopic())
	if err != nil {
		return nil, err
	}

	client, err := cloudevents.NewClient(sender, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		return nil, err
	}

	return &KafkaCloudEventsProducer{
		client: client,
		source: config.EventSource,
	}, nil
}

func (c *KafkaCloudEventsProducer) Send(ctx context.Context, event Event) error {
	ce, err := toCloudEvent(event, c.source)
	if err != nil {
		return err
	}

	if key := event.GetEventKey(); key != "" {
		ctx = kafka_sarama.WithMessageKey(ctx, sarama.StringEncoder(key))
	}
	return c.SendCloudEvent(ctx, ce)
}

func (c *KafkaCloudEventsProducer) SendCloudEvent(ctx context.Context, event cloudevents.Event) error {
	return c.client.Send(
		cloudevents.ContextWithRetriesExponentialBackoff(ctx, producerRetryPeriod,producerMaxRetries),
		event,
	)
}

func toCloudEvent(event Event, source string) (cloudevents.Event, error) {
	e := cloudevents.NewEvent()
	e.SetType(event.GetEventType())
	e.SetSource(source)
	if err := e.SetData(cloudevents.ApplicationJSON, event); err != nil {
		return e, err
	}

	return e, nil
}
