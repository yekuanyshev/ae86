package transport

import (
	"github.com/supernova0730/ae86/internal/transport/rest"
)

type Config struct {
	Rest rest.Config
}

func Start(config Config) {
	go rest.Start(config.Rest)
}
