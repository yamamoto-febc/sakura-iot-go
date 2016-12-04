package sakura_iot_go

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var payloadTestJSONTemplate = `{
    "module": "XXXXXXXXX",
    "type": "channels",
    "datetime": "2016-06-01T12:21:11.628907163Z",
    "payload": {
        "channels": [
                 %s
        ]
    }
}`

var channelTestJSONInt = `{
    "channel": 1,
    "type": "i",
    "value": 1
}`

var channelTestJSONHexString = `{
    "channel": 2,
    "type": "b",
    "value": "0f1e2d3c4b5c6b7a"
}`

var keepAliveTestJSON = `{"type": "keepalive", "datetime": "2016-06-11T06:24:50.643930807Z"}`

var payloadTestJSONInt = fmt.Sprintf(
	payloadTestJSONTemplate,
	channelTestJSONInt,
)
var payloadTestJSONHexString = fmt.Sprintf(
	payloadTestJSONTemplate,
	channelTestJSONHexString,
)

var payloadTestJSONChannelArray = fmt.Sprintf(
	payloadTestJSONTemplate,
	fmt.Sprintf("%s,%s", channelTestJSONInt, channelTestJSONHexString),
)

func TestPayloadUnmarshalJSONBasic(t *testing.T) {

	var payload Payload
	err := json.Unmarshal([]byte(payloadTestJSONInt), &payload)

	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.Equal(t, payload.Module, "XXXXXXXXX")
	assert.Equal(t, payload.Type, PayloadTypesChannels)
	//2016-06-01T12:21:11.628907163Z
	assert.Equal(t, payload.Datetime.Year(), 2016)
	assert.Equal(t, payload.Datetime.Month(), time.Month(6))
	assert.Equal(t, payload.Datetime.Day(), 1)

	// do test in JST
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", int(9*time.Hour))
	}
	jst := payload.Datetime.In(loc)

	assert.Equal(t, jst.Hour(), 21)
	assert.Equal(t, jst.Minute(), 21)
	assert.Equal(t, jst.Second(), 11)
	assert.Equal(t, jst.Nanosecond(), 628907163)

}

func TestPayloadUnmarshalJSONIntChannels(t *testing.T) {

	var payload Payload
	err := json.Unmarshal([]byte(payloadTestJSONInt), &payload)

	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.NotNil(t, payload.Payload)
	assert.Len(t, payload.Payload.Channels, 1)

	channelValue := payload.Payload.Channels[0]

	assert.Equal(t, channelValue.Type, "i")

	intValue, err := channelValue.GetInt()
	assert.NoError(t, err)
	assert.Equal(t, intValue, int32(1))

	strValue, err := channelValue.GetHexString()
	assert.Error(t, err)
	assert.Equal(t, strValue, "")

}

func TestPayloadUnmarshalJSONHexStringChannels(t *testing.T) {

	var payload Payload
	err := json.Unmarshal([]byte(payloadTestJSONHexString), &payload)

	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.NotNil(t, payload.Payload)
	assert.Len(t, payload.Payload.Channels, 1)

	channelValue := payload.Payload.Channels[0]

	assert.Equal(t, channelValue.Type, "b")

	intValue, err := channelValue.GetInt()
	assert.Error(t, err)
	assert.Equal(t, intValue, int32(0))

	strValue, err := channelValue.GetHexString()
	assert.NoError(t, err)
	assert.Equal(t, strValue, "0f1e2d3c4b5c6b7a")

}

func TestPayloadUnmarshalJSONWithArray(t *testing.T) {
	var payload Payload
	err := json.Unmarshal([]byte(payloadTestJSONChannelArray), &payload)

	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.Len(t, payload.Payload.Channels, 2)

}

func TestPayloadUnmarshalJSONKeepAlive(t *testing.T) {

	var payload Payload
	err := json.Unmarshal([]byte(keepAliveTestJSON), &payload)

	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.True(t, payload.IsKeepAlive())
	assert.False(t, payload.IsChannelValue())

}
