package whatsapp

import (
	"github.com/probnotas/nanoClaw/pkg/bus"
	"github.com/probnotas/nanoClaw/pkg/channels"
	"github.com/probnotas/nanoClaw/pkg/config"
)

func init() {
	channels.RegisterFactory("whatsapp", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewWhatsAppChannel(cfg.Channels.WhatsApp, b)
	})
}
