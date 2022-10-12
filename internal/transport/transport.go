package transport

import (
	"github.com/supernova0730/ae86/internal/container"
	"github.com/supernova0730/ae86/internal/transport/rest"
)

type Config struct {
	Rest rest.Config
}

func Start(config Config, controller *container.ControllerContainer) {
	go rest.Start(config.Rest, controller)
}
