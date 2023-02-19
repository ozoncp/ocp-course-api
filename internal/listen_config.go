package internal

import (
	"fmt"
	"net"

	"github.com/go-akka/configuration"

	uc "github.com/ozoncp/ocp-course-api/internal/utils/config"
)

type ListenConfig struct {
	Interface string
	Port      int
}

func (c *ListenConfig) Validate() error {

	errs := make([]string, 0, 2)

	if net.ParseIP(c.Interface) == nil {
		errs = append(errs,
			fmt.Sprintf("the Interface value is invalid, %v is bad IP", c.Interface))
	}

	if p := c.Port; p < 1 || p > 65535 {
		errs = append(errs,
			fmt.Sprintf("the Port value is invalid, %v is bad", p))
	}

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

func (this *ListenConfig) Address() string {
	return fmt.Sprintf("%v:%v", this.Interface, this.Port)
}

func FromHoconListenConfig(cfg *configuration.Config, path string) (config *ListenConfig, err error) {
	defer func() {
		if exception := recover(); exception != nil {
			config = nil
			if errIn, ok := exception.(error); ok {
				err = fmt.Errorf("reading config failed: %w", errIn)
			} else {
				err = fmt.Errorf("reading config failed: %v", exception)
			}
		}
	}()

	if path != "" {
		if cfg, err := uc.GetConfig(cfg, path); err != nil {
			return nil, err
		} else {
			return FromHoconListenConfig(cfg, "")
		}
	}

	listenInterface := cfg.GetString("interface")
	port := int(cfg.GetInt32("port"))

	config = &ListenConfig{Interface: listenInterface, Port: port}

	if err = config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}
