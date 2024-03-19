package option

import (
	"fmt"
)

type config struct {
	host string
	port uint
}

func new(opts ...Option[config]) *config {
	config := config{}

	if opts != nil {
		Apply(&config, opts...)
	}

	return &config
}

func withHost(server string) Option[config] {
	return func(cfg *config) error {
		cfg.host = server
		return nil
	}
}

func withPort(port uint) Option[config] {
	return func(cfg *config) error {
		cfg.port = port
		return nil
	}
}

func ExampleApply() {
	cfg := new()
	Apply(cfg, withHost("localhost"))
	fmt.Println(cfg.host)
	// Output: localhost
}

func ExampleGroup() {
	grp := Group(
		withHost("localhost"),
		withPort(8000),
	)
	cfg := new(grp)
	fmt.Println(cfg.host, cfg.port)
	// Output: localhost 8000
}
