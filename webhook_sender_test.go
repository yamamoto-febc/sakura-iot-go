package sakura_iot_go

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWebhookSender_Send(t *testing.T) {

	token := os.Getenv("SAKURA_IOT_SENDER_TOKEN")
	secret := os.Getenv("SAKURA_IOT_SENDER_SECRET")
	module := os.Getenv("SAKURA_IOT_SENDER_MODULE")

	if token == "" {
		t.Logf("WebhookSender:token is required. skip.")
		t.SkipNow()
		return
	}
	if module == "" {
		t.Logf("WebhookSender:module is required. skip.")
		t.SkipNow()
		return

	}

	sender := NewWebhookSender(token, secret)
	p := NewPayload(module)
	p.AddChannelByInt(0, -1)
	p.AddChannelByInt(1, 0)
	p.AddChannelByInt(2, 1)
	p.AddChannelByInt(3, 255)
	p.AddChannelByInt(4, 256)
	p.AddChannelByInt(5, 257)

	err := sender.Send(p)
	assert.NoError(t, err)

}
