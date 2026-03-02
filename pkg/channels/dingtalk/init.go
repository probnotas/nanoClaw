package dingtalk

import (
	"github.com/probnotas/nanoClaw/pkg/bus"
	"github.com/probnotas/nanoClaw/pkg/channels"
	"github.com/probnotas/nanoClaw/pkg/config"
)

func init() {
	channels.RegisterFactory("dingtalk", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewDingTalkChannel(cfg.Channels.DingTalk, b)
	})
}
