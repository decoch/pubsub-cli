package action

type Topic struct {
	ID            string         `yaml:"topic_id"`
	Subscriptions []Subscription `yaml:"subscriptions"`
}

type Subscription struct {
	ID       string `yaml:"id"`
	Endpoint string `yaml:"endpoint"`
}
