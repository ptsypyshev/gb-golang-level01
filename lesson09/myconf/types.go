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

//type MyConfigJSON struct {
//	Port        int    `json:"port"`
//	DbUrl       string `json:"db_url"`
//	JaegerUrl   string `json:"jaeger_url"`
//	SentryUrl   string `json:"sentry_url"`
//	KafkaBroker string `json:"kafka_broker"`
//	SomeAppId   int    `json:"some_app_id"`
//	SomeAppKey  string `json:"some_app_key"`
//}
//
//type MyConfigYAML struct {
//	Port        int    `yaml:"port"`
//	DbUrl       string `yaml:"db_url"`
//	JaegerUrl   string `yaml:"jaeger_url"`
//	SentryUrl   string `yaml:"sentry_url"`
//	KafkaBroker string `yaml:"kafka_broker"`
//	SomeAppId   int    `yaml:"some_app_id"`
//	SomeAppKey  string `yaml:"some_app_key"`
//}

type MyConfigFile struct {
	Port        int    `json:"port" yaml:"port"`
	DbUrl       string `json:"db_url" yaml:"db_url"`
	JaegerUrl   string `json:"jaeger_url" yaml:"jaeger_url"`
	SentryUrl   string `json:"sentry_url" yaml:"sentry_url"`
	KafkaBroker string `json:"kafka_broker" yaml:"kafka_broker"`
	SomeAppId   int    `json:"some_app_id" yaml:"some_app_id"`
	SomeAppKey  string `json:"some_app_key" yaml:"some_app_key"`
}
