package feishu

import (
	"github.com/probnotas/nanoClaw/pkg/bus"
	"github.com/probnotas/nanoClaw/pkg/channels"
	"github.com/probnotas/nanoClaw/pkg/config"
)

func init() {
	channels.RegisterFactory("feishu", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewFeishuChannel(cfg.Channels.Feishu, b)
	})
}
