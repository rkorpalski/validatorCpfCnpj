package env

import (
	"os"
	"strconv"
	"strings"
)

var properties = []*Property{
	{Key:"APP_BASE_URL", FallBackValue:"0.0.0.0:8080"},
	{Key:"MONGO_BASE_URL", FallBackValue:"mongodb://mongo:27017"},
	{Key:"APP_DATABASE", FallBackValue:"cpfCnpj"},
	{Key:"API_CONTEXT", FallBackValue:"/validator"},
}

type Property struct {
	Key 			string
	FallBackValue 	string
}

type Properties struct {
	Properties []*Property
}

func envProperties() *Properties {
	objProperties := &Properties{}
	objProperties.Properties = properties
	return objProperties
}

func (e *Properties) getProperty(env string) *Property {
	for _, property := range properties {
		if strings.ToUpper(property.Key) == strings.ToUpper(env) {
			return property
		}
	}
	return nil
}

func (e *Properties) envInt(env string) int  {
	b, _ := strconv.Atoi(e.envString(env))
	return b
}

func (e * Properties) envString(env string) string {
	property := e.getProperty(env)
	if property != nil {
		value := os.Getenv(env)
		if value == "" {
			return property.FallBackValue
		}
		return value
	}
	return ""
}

func AppBaseUrl() string {
	return envProperties().envString("APP_BASE_URL")
}

func MongoBaseUrl() string {
	return envProperties().envString("MONGO_BASE_URL")
}

func AppDatabase() string {
	return envProperties().envString("APP_DATABASE")
}

func ApiContext() string {
	return envProperties().envString("API_CONTEXT")
}