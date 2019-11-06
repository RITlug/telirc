package telegram

import (
	"testing"

	"github.com/ritlug/teleirc/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewClientBasic(t *testing.T) {
	tgSettings := internal.TelegramSettings{
		Token:  "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatID: "-0000000000000",
	}
	client := NewClient(tgSettings)
	assert.Equal(t, client.Settings, tgSettings, "Basic client settings should be properly set")
}

func TestNewClientFull(t *testing.T) {
	tgSettings := internal.TelegramSettings{
		Token:               "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatID:              "-0000000000000",
		ShowJoinMessage:     true,
		ShowActionMessage:   true,
		ShowLeaveMessage:    true,
		ShowKickMessage:     true,
		MaxMessagePerMinute: 10,
	}
	client := NewClient(tgSettings)
	assert.Equal(t, client.Settings, tgSettings, "All client settings should be properly set")
}