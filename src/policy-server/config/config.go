package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	validator "gopkg.in/validator.v2"

	"code.cloudfoundry.org/cf-networking-helpers/db"
)

type Config struct {
	ListenHost                      string    `json:"listen_host" validate:"nonzero"`
	ListenPort                      int       `json:"listen_port" validate:"nonzero"`
	LogPrefix                       string    `json:"log_prefix" validate:"nonzero"`
	DebugServerHost                 string    `json:"debug_server_host" validate:"nonzero"`
	DebugServerPort                 int       `json:"debug_server_port" validate:"nonzero"`
	UAAClient                       string    `json:"uaa_client" validate:"nonzero"`
	UAAClientSecret                 string    `json:"uaa_client_secret" validate:"nonzero"`
	UAACA                           string    `json:"uaa_ca"`
	UAAURL                          string    `json:"uaa_url" validate:"nonzero"`
	UAAPort                         int       `json:"uaa_port" validate:"nonzero"`
	CCURL                           string    `json:"cc_url" validate:"nonzero"`
	SkipSSLValidation               bool      `json:"skip_ssl_validation"`
	Database                        db.Config `json:"database" validate:"nonzero"`
	TagLength                       int       `json:"tag_length" validate:"nonzero"`
	MetronAddress                   string    `json:"metron_address" validate:"nonzero"`
	LogLevel                        string    `json:"log_level"`
	CleanupInterval                 int       `json:"cleanup_interval" validate:"min=1"`
	CCAppRequestChunkSize           int       `json:"cc_app_request_chunk_size"`
	RequestTimeout                  int       `json:"request_timeout" validate:"min=1"`
	MaxPolicies                     int       `json:"max_policies" validate:"min=1"`
	EnableSpaceDeveloperSelfService bool      `json:"enable_space_developer_self_service"`
	AllowedCORSDomains              []string  `json:"allowed_cors_domains"`
}

func (c *Config) Validate() error {
	return validator.Validate(c)
}

func New(path string) (*Config, error) {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %s", err)
	}

	cfg := Config{}
	err = json.Unmarshal(jsonBytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %s", err)
	}

	if err := cfg.Validate(); err != nil {
		return &cfg, fmt.Errorf("invalid config: %s", err)
	}

	return &cfg, nil
}
