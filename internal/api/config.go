package api

import (
	"fmt"
	"os"

	"github.com/go-akka/configuration"

	"github.com/ozoncp/ocp-course-api/internal"
	uc "github.com/ozoncp/ocp-course-api/internal/utils/config"
)

type Config struct {
	Grpc        internal.ListenConfig
	Http        internal.ListenConfig
	SwaggerFile string
}

func NewConfig(listenInterface string, grpcPort int, httpPort int, swagger string) *Config {
	return &Config{
		Grpc:        internal.ListenConfig{Interface: listenInterface, Port: grpcPort},
		Http:        internal.ListenConfig{Interface: listenInterface, Port: httpPort},
		SwaggerFile: swagger,
	}
}

func (c *Config) Validate() error {
	if c == nil {
		return nil
	}

	errs := make([]string, 0, 3)

	if e := c.Grpc.Validate(); e != nil {
		errs = append(errs, fmt.Sprintf("the Grpc is invalid: [%v]", e))
	}

	if e := c.Http.Validate(); e != nil {
		errs = append(errs, fmt.Sprintf("the Http is invalid: [%v]", e))
	}

	func() {
		stat, e := os.Stat(c.SwaggerFile)

		if e != nil {
			errs = append(errs,
				fmt.Sprintf("the SwaggerFile is invalid: '%v' doesn't exists", c.SwaggerFile))
			return
		}

		if stat.IsDir() {
			errs = append(errs,
				fmt.Sprintf("the SwaggerFile is invalid: '%v' is a directory", c.SwaggerFile))
			return
		}
	}()

	err := ""
	for _, e := range errs {
		if len(err) > 0 {
			err = err + "; " + e
		} else {
			err = e
		}
	}

	if len(err) > 0 {
		return fmt.Errorf(err)
	}

	return nil
}

func FromHoconConfig(cfg *configuration.Config, path string) (config *Config, err error) {
	defer func() {
		if errIn := recover(); errIn != nil {
			config = nil
			err = fmt.Errorf("reading config failed: %w", errIn)
		}
	}()

	if path != "" {
		if cfg, err := uc.GetConfig(cfg, path); err != nil {
			return nil, err
		} else {
			return FromHoconConfig(cfg, "")
		}
	}

	listenInterface := cfg.GetString("interface")
	grpcPort := int(cfg.GetInt32("grpc-port"))
	httpPort := int(cfg.GetInt32("http-port"))
	swaggerFile := cfg.GetString("swagger-file")

	config = NewConfig(listenInterface, grpcPort, httpPort, swaggerFile)

	if err = config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}
