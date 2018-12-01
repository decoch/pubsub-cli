package action

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"cloud.google.com/go/pubsub"
	"gopkg.in/yaml.v2"
)

// Create subscription from yaml file.
func Create(ctx context.Context, projectID, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var topics []Topic
	err = yaml.Unmarshal(data, &topics)
	if err != nil {
		return err
	}

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		t := client.Topic(topic.ID)
		exists, err := t.Exists(ctx)
		if err != nil {
			return err
		}
		if !exists {
			t, err = client.CreateTopic(ctx, topic.ID)
			if err != nil {
				return err
			}
		}

		for _, sub := range topic.Subscriptions {
			_, err = client.CreateSubscription(ctx, sub.ID, pubsub.SubscriptionConfig{
				Topic:       t,
				AckDeadline: 10 * time.Second,
				PushConfig:  pubsub.PushConfig{Endpoint: sub.Endpoint},
			})

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}
