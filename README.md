Still developing

# pubsub-cli

[![GoDoc](https://godoc.org/github.com/decoch/pubsub-cli?status.svg)](https://godoc.org/github.com/decoch/pubsub-cli)

pubsub-cli is a command line tool for Google Cloud Pub/Sub.

Easy to create subscriptions can be confirmed.

## Features

- Get list of subscription.
- Create subscription.

## Installation

```bash
go get -u github.com/decoch/pubsub-cli
```

## How to use

```bash
$ pubsub-cli list project-id
$ pubsub-cli create project-id filename.yaml
```

yaml format

```yaml
- topic_id: analytics
  subscriptions:
  - id: lead_tag
    endpoint: http://localhost:8080/leads/tags
- topic_id: analytics2
  subscriptions:
  - id: lead_tag2
    endpoint: http://localhost:8080/leads/tags2
```