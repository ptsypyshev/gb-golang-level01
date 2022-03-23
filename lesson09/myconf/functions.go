package myconf

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
	"strconv"
)

const (
	C_JSON = iota
	C_YAML
)

func GetFromFile(f string, t int) (MyConfig MyConfigFile) {
	var errUnmarshal error
	configFile, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	switch t {
	case C_JSON:
		errUnmarshal = json.Unmarshal(configFile, &MyConfig)
	case C_YAML:
		errUnmarshal = yaml.Unmarshal(configFile, &MyConfig)
	}
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}
	return
}

func GetFromEnvs() map[string]string {
	param := make(map[string]string)
	for _, v := range Parameters {
		param[v] = os.Getenv(v)
	}
	return param
}

func GetFromFlags() map[string]string {
	param := make(map[string]string)
	port := flag.String("port", "8080", "Which port should application listen")
	//port := flag.Int("port", 8080, "Which port should application listen")
	dbUrlFlag := flag.String("db_url", "", "Specify Database URL (example, "+
		"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable)")
	jaegerUrlFlag := flag.String("jaeger_url", "", "Specify Jaeger URL (example, "+
		"http://jaeger:16686)")
	sentryUrlFlag := flag.String("sentry_url", "", "Specify Sentry URL (example, "+
		"http://sentry:9000)")
	kafkaBroker := flag.String("kafka_broker", "", "Specify Kafka URL (example, "+
		"kafka:9092)")
	someAppId := flag.String("some_app_id", "0", "Specify Application ID")
	//someAppId := flag.Int("some_app_id", 0, "Specify Application ID")
	someAppKey := flag.String("some_app_key", "", "Specify Application Key (example, "+
		"Idgz1PE3zO9iNc0E3oeH3CHDPX9MzZe3)")
	jsonFile := flag.String("json", "", "Specify Full Path to JSON file (example, "+
		"/home/user/config.jsonFile)")
	yamlFile := flag.String("yaml", "", "Specify Full Path to YAML file (example, "+
		"/home/user/config.yamlFile)")
	flag.Parse()

	if *jsonFile != "" {
		param["json"] = *jsonFile
		return param
	}
	if *yamlFile != "" {
		param["yaml"] = *yamlFile
		return param
	}

	for _, v := range Parameters {
		switch v {
		case "port":
			param[v] = *port
		case "db_url":
			param[v] = *dbUrlFlag
		case "jaeger_url":
			param[v] = *jaegerUrlFlag
		case "sentry_url":
			param[v] = *sentryUrlFlag
		case "kafka_broker":
			param[v] = *kafkaBroker
		case "some_app_id":
			param[v] = *someAppId
		case "some_app_key":
			param[v] = *someAppKey
		}
	}
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

func ParseMap(p map[string]string) (MyConfig, error) {
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

func (m MyConfigFile) ParseToMyConfig() (res MyConfig, err error) {
	var errorSl [5]error
	res.Port = m.Port
	res.SomeAppId = m.SomeAppId
	res.SomeAppKey, errorSl[0] = ParseString("SomeAppKey", m.SomeAppKey)
	res.KafkaBroker, errorSl[1] = ParseString("KafkaBroker", m.KafkaBroker)
	res.DbUrl, errorSl[2] = ParseUrl("DbUrl", m.DbUrl)
	res.JaegerUrl, errorSl[3] = ParseUrl("JaegerUrl", m.JaegerUrl)
	res.SentryUrl, errorSl[4] = ParseUrl("SentryUrl", m.SentryUrl)
	for _, e := range errorSl {
		if e != nil {
			return res, fmt.Errorf("cannot parse from config file: %w", e)
		}
	}
	return res, nil
}

func PrintDef() {
	if flag.NFlag() > 1 {
		fmt.Println("You should specify all next arguments:")
		flag.PrintDefaults()
	}
}
