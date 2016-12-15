package sakura

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
	p.AddValueByInt(0, -1)
	p.AddValueByInt(1, 0)
	p.AddValueByInt(2, 1)
	p.AddValueByInt(3, 255)
	p.AddValueByInt(4, 256)
	p.AddValueByInt(5, 257)

	err := sender.Send(p)
	assert.NoError(t, err)

}
