package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	defaultPrefix string = "+"
)

type Config struct {
	Bot      bool
	Token    string
	Channels []string
	Prefix   string
}

func (c *Config) IsChannel(s string) bool {
	for _, channel := range c.Channels {
		if channel == s {
			return true
		}
	}

	return false
}

func LoadConfig(c string) (Config, error) {
	file, err := os.Open(c)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return Config{}, err
	}

	if config.Bot {
		config.Token = "Bot " + config.Token
	}

	channels := []string{}

	for _, channel := range config.Channels {
		if channel == "" || len(channel) < 1 {
			continue
		}

		channels = append(channels, channel)
	}

	prefix := defaultPrefix
	if config.Prefix != "" {
		prefix = config.Prefix
	}

	return Config{config.Bot, config.Token, channels, prefix}, nil
}
