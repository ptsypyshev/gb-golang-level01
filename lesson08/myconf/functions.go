package myconf

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func GetFromEnvs() map[string]string {
	param := make(map[string]string)
	for _, v := range Parameters {
		param[v] = os.Getenv(v)
	}
	return param
}

func GetFromFlags() map[string]string {
	var port, dbUrlFlag, jaegerUrlFlag, sentryUrlFlag, kafkaBroker, someAppId, someAppKey string
	param := make(map[string]string)

	flag.StringVar(&port, "port", "8080", "Which port should application listen")
	flag.StringVar(&dbUrlFlag, "db_url", "", "Specify Database URL (example, "+
		"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable)")
	flag.StringVar(&jaegerUrlFlag, "jaeger_url", "", "Specify Jaeger URL (example, "+
		"http://jaeger:16686)")
	flag.StringVar(&sentryUrlFlag, "sentry_url", "", "Specify Sentry URL (example, "+
		"http://sentry:9000)")
	flag.StringVar(&kafkaBroker, "kafka_broker", "", "Specify Kafka URL (example, "+
		"kafka:9092)")
	flag.StringVar(&someAppId, "some_app_id", "0", "Specify Application ID")
	flag.StringVar(&someAppKey, "some_app_key", "", "Specify Application Key (example, "+
		"Idgz1PE3zO9iNc0E3oeH3CHDPX9MzZe3)")

	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		param[f.Name] = f.Value.String()
	})

	return param
}

func ParseInt(k, v string) (int, error) {
	intVal, err := strconv.Atoi(v)
	if err != nil {
		PrintDef()
		return 0, fmt.Errorf("cannot parse %s value %s to integer", k, v)
	}
	return intVal, nil
}

func ParseUrl(k, v string) (*url.URL, error) {
	parsedUrl, err := url.ParseRequestURI(v)
	if err != nil {
		PrintDef()
		return nil, fmt.Errorf("cannot parse %s value %s to URL", k, v)
	}
	return parsedUrl, nil
}

func ParseString(k, v string) (string, error) {
	if v == "" {
		PrintDef()
		err := fmt.Errorf("empty %s value", k)
		return v, err
	}
	return v, nil
}

func Parser(p map[string]string) (MyConfig, error) {
	fmt.Println(p)
	result := MyConfig{}
	var err error
	for k, v := range p {
		switch k {
		case "port":
			result.Port, err = ParseInt(k, v)
		case "db_url":
			result.DbUrl, err = ParseUrl(k, v)
		case "jaeger_url":
			result.JaegerUrl, err = ParseUrl(k, v)
		case "sentry_url":
			result.SentryUrl, err = ParseUrl(k, v)
		case "kafka_broker":
			result.KafkaBroker, err = ParseString(k, v)
		case "some_app_id":
			result.SomeAppId, err = ParseInt(k, v)
		case "some_app_key":
			result.SomeAppKey, err = ParseString(k, v)
		}
		if err != nil {
			return MyConfig{}, err
		}
	}
	return result, nil
}

func PrintDef() {
	fmt.Println("You should specify all next arguments:")
	flag.PrintDefaults()
}
