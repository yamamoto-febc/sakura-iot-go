package sakura

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/yamamoto-febc/sakura-iot-go/version"
	"io/ioutil"
	"net/http"
)

// WebhookSendRootURL is URL prefix of send webhook target
var WebhookSendRootURL = "https://api.sakura.io/incoming/v1/"

// WebhookSenderUserAgent user-agent string
var WebhookSenderUserAgent = fmt.Sprintf("sakura-iot-go/%s", version.Version)

// WebhookSender is type to handling Webhook that send to Sakura-IoT-platform
type WebhookSender struct {
	Token  string
	Secret string
}

// NewWebhookSender create new *WebhookSender
func NewWebhookSender(token string, secret string) *WebhookSender {
	return &WebhookSender{
		Token:  token,
		Secret: secret,
	}
}

// Send send new request to the Incoming-Webhook on Sakura-IoT-platform
func (w *WebhookSender) Send(p Payload) error {
	var (
		client = &http.Client{}
		url    = fmt.Sprintf("%s/%s", WebhookSendRootURL, w.Token)
		err    error
		req    *http.Request
	)

	var bodyJSON []byte
	bodyJSON, err = json.Marshal(p)
	if err != nil {
		return fmt.Errorf("Failed on Marshaling payload : %s", err)
	}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return fmt.Errorf("Failed on creating new request: %s", err)
	}

	req.Header.Add("User-Agent", WebhookSenderUserAgent)
	req.Header.Add("Content-Type", "application/json")

	if w.Secret != "" {

		computed := hmac.New(sha1.New, []byte(w.Secret))
		computed.Write(bodyJSON)
		signBody := []byte(computed.Sum(nil))

		req.Header.Add("X-Sakura-Signature", hex.EncodeToString(signBody))
	}

	req.Method = "POST"
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Faild on sending request:%s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	// here, on error
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("Send webhook failed:%s", string(data))

}
