package pico

import (
	"github.com/probnotas/nanoClaw/pkg/bus"
	"github.com/probnotas/nanoClaw/pkg/channels"
	"github.com/probnotas/nanoClaw/pkg/config"
)

func init() {
	channels.RegisterFactory("nano", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewNanoChannel(cfg.Channels.Pico, b)
	})
}
