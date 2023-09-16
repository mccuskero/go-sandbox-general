package config

import (
	"errors"
)

type Config struct {
	AwsRegion string
}

func NewFromRegion(region string) (*Config, error) {

	if region == "" {
		return nil, errors.New("region is empty")
	}

	cfg := &Config{
		AwsRegion: region,
	}

	return cfg, nil
}
