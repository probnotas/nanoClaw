package onebot

import (
	"github.com/probnotas/nanoClaw/pkg/bus"
	"github.com/probnotas/nanoClaw/pkg/channels"
	"github.com/probnotas/nanoClaw/pkg/config"
)

func init() {
	channels.RegisterFactory("onebot", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewOneBotChannel(cfg.Channels.OneBot, b)
	})
}
