package action

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
	"gopkg.in/yaml.v2"
)

// List subscriptions from pub/sub.
func List(ctx context.Context, projectID, filename string) error {
	topics, err := getSubscriptions(ctx, projectID)
	if err != nil {
		return err
	}

	if filename == "" {
		show(topics)
		return nil
	}
	return write(filename, topics)
}

func getSubscriptions(ctx context.Context, projectID string) ([]Topic, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	defer client.Close()

	if err != nil {
		return []Topic{}, err
	}

	var topics []Topic
	it := client.Topics(ctx)
	for {
		t, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []Topic{}, err
		}

		var topic Topic
		topic.ID = t.ID()
		subs := t.Subscriptions(ctx)
		for {
			s, err := subs.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return []Topic{}, err
			}

			config, err := s.Config(ctx)
			if err != nil {
				return []Topic{}, err
			}

			subscription := Subscription{
				ID:       s.ID(),
				Endpoint: config.PushConfig.Endpoint,
			}

			topic.Subscriptions = append(topic.Subscriptions, subscription)
		}

		topics = append(topics, topic)
	}
	return topics, nil
}

func show(topics []Topic) {
	for _, topic := range topics {
		fmt.Printf("- topic_id: %s\n", topic.ID)
		if len(topic.Subscriptions) > 0 {
			fmt.Printf("  subscriptions: %s\n", topic.ID)
		}
		for _, sub := range topic.Subscriptions {
			fmt.Printf("  - id: %s\n", sub.ID)
			fmt.Printf("    endpoint: %s\n", sub.Endpoint)
		}
	}
}

func write(filename string, topics []Topic) error {
	data, err := yaml.Marshal(topics)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
