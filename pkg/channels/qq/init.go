package qq

import (
	"github.com/probnotas/nanoClaw/pkg/bus"
	"github.com/probnotas/nanoClaw/pkg/channels"
	"github.com/probnotas/nanoClaw/pkg/config"
)

func init() {
	channels.RegisterFactory("qq", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewQQChannel(cfg.Channels.QQ, b)
	})
}
