package main

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	*ApplicationConfig
}

type ApplicationConfig struct {
	Server    ServerConfig    `yaml:"server"`
	Couchbase CouchbaseConfig `yaml:"couchbase"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type CouchbaseConfig struct {
	Addresses string `yaml:"addresses"`
	Bucket    string `yaml:"bucket"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

func NewConfiguration() *Config {
	applicationConfig := &ApplicationConfig{}
	applicationConfig.readApplicationConfig()

	return &Config{ApplicationConfig: applicationConfig}
}

func (c *ApplicationConfig) readApplicationConfig() {
	env, found := os.LookupEnv("ACTIVE_PROFILE")

	if !found {
		env = "local"

	}

	print("ACTIVE_PROFILE: ", env, "\n")

	v := viper.New()

	v.SetTypeByDefaultValue(true)
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath("./resource")

	readConfigErr := v.ReadInConfig()

	if readConfigErr != nil {
		panic("Couldn't load application configuration, cannot start. Terminating. : " + readConfigErr.Error())
	}

	sub := v.Sub(env)

	unMarshallErr := sub.Unmarshal(c)

	if unMarshallErr != nil {
		panic("Configuration cannot deserialize. Terminating. : " + unMarshallErr.Error())
	}
}
