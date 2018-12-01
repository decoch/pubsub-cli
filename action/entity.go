package action

type Topic struct {
	TopicID       string         `yaml:"topic"`
	Subscriptions []Subscription `yaml:"subscriptions"`
}

type Subscription struct {
	SubscriptionID       string `yaml:"subscription"`
	SubscriptionEndpoint string `yaml:"endpoint"`
}
