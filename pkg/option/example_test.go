package option

import (
	"fmt"
)

type Config struct {
	host string
	port uint
}

func New(opts ...Option[Config]) *Config {
	config := Config{}

	if opts != nil {
		Apply(&config, opts...)
	}

	return &config
}

func WithHost(server string) Option[Config] {
	return func(cfg *Config) error {
		cfg.host = server
		return nil
	}
}

func WithPort(port uint) Option[Config] {
	return func(cfg *Config) error {
		cfg.port = port
		return nil
	}
}

func ExampleApply() {
	cfg := New()
	Apply(cfg, WithHost("localhost"))
	fmt.Println(cfg.host)
	// Output: localhost
}

func ExampleGroup() {
	grp := Group(
		WithHost("localhost"),
		WithPort(8000),
	)
	cfg := New(grp)
	fmt.Println(cfg.host, cfg.port)
	// Output: localhost 8000
}
