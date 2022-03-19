package myconf

import "net/url"

var Parameters = [...]string{"port", "db_url", "jaeger_url", "sentry_url", "kafka_broker", "some_app_id", "some_app_key"}

type MyConfig struct {
	Port        int
	DbUrl       *url.URL
	JaegerUrl   *url.URL
	SentryUrl   *url.URL
	KafkaBroker string
	SomeAppId   int
	SomeAppKey  string
}
