package gofer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/hcl/v2"

	ethereumConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/ethereum"
	loggerConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/logger"
	priceProviderConfig "github.com/chronicleprotocol/oracle-suite/pkg/config/priceprovidernext"
	"github.com/chronicleprotocol/oracle-suite/pkg/log"
	"github.com/chronicleprotocol/oracle-suite/pkg/pricenext/provider"
	pkgSupervisor "github.com/chronicleprotocol/oracle-suite/pkg/supervisor"
)

// Config is the configuration for Gofer.
type Config struct {
	Gofer    priceProviderConfig.Config `hcl:"gofernext,block"`
	Ethereum *ethereumConfig.Config     `hcl:"ethereum,block,optional"`
	Logger   *loggerConfig.Config       `hcl:"logger,block,optional"`

	// HCL fields:
	Remain  hcl.Body        `hcl:",remain"` // To ignore unknown blocks.
	Content hcl.BodyContent `hcl:",content"`
}

// Services returns the services that are configured from the Config struct.
type Services struct {
	PriceProvider provider.Provider
	Logger        log.Logger

	supervisor *pkgSupervisor.Supervisor
}

// Start implements the supervisor.Service interface.
func (s *Services) Start(ctx context.Context) error {
	if s.supervisor != nil {
		return fmt.Errorf("services already started")
	}
	s.supervisor = pkgSupervisor.New(s.Logger)
	if p, ok := s.PriceProvider.(pkgSupervisor.Service); ok {
		s.supervisor.Watch(p)
	}
	if l, ok := s.Logger.(pkgSupervisor.Service); ok {
		s.supervisor.Watch(l)
	}
	return s.supervisor.Start(ctx)
}

// Wait implements the supervisor.Service interface.
func (s *Services) Wait() <-chan error {
	return s.supervisor.Wait()
}

// Services returns the services configured for Gofer.
func (c *Config) Services(baseLogger log.Logger) (*Services, error) {
	logger, err := c.Logger.Logger(loggerConfig.Dependencies{
		AppName:    "gofer",
		BaseLogger: baseLogger,
	})
	if err != nil {
		return nil, err
	}
	clients, err := c.Ethereum.ClientRegistry(ethereumConfig.Dependencies{Logger: logger})
	if err != nil {
		return nil, err
	}
	priceProvider, err := c.Gofer.PriceProvider(priceProviderConfig.Dependencies{
		HTTPClient: &http.Client{},
		Clients:    clients,
		Logger:     logger,
	})
	if err != nil {
		return nil, err
	}
	return &Services{
		PriceProvider: priceProvider,
		Logger:        logger,
	}, nil
}
