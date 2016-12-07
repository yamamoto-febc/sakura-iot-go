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
    "value": 1,
    "datetime":"2016-12-04T04:14:27.214224349Z"
}`

var channelTestJSONHexString = `{
    "channel": 2,
    "type": "b",
    "value": "0f1e2d3c4b5c6b7a",
    "datetime":"2016-12-04T04:14:27.214224349Z"
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

func TestPayloadHandleChannels(t *testing.T) {
	payload := NewPayload("xxxxxxxx10xx")

	assert.NotNil(t, payload)
	assert.Equal(t, payload.Module, "xxxxxxxx10xx")
	assert.NotNil(t, payload.Payload)
	assert.Len(t, payload.Payload.Channels, 0)

	// hexString
	payload.AddChannelByHexString(0, "FF01FF01")

	assert.Len(t, payload.Payload.Channels, 1)
	assert.EqualValues(t, payload.Payload.Channels[0].Channel, 0)
	assert.Equal(t, payload.Payload.Channels[0].Type, "b")
	assert.EqualValues(t, payload.Payload.Channels[0].Value, "FF01FF01")

	// int32
	payload.AddChannelByInt(1, 1)

	assert.Len(t, payload.Payload.Channels, 2)
	assert.EqualValues(t, payload.Payload.Channels[1].Channel, 1)
	assert.Equal(t, payload.Payload.Channels[1].Type, "i")
	assert.EqualValues(t, payload.Payload.Channels[1].Value, 1)

	// uint32
	payload.AddChannelByUint(2, uint32(1))

	assert.Len(t, payload.Payload.Channels, 3)
	assert.EqualValues(t, payload.Payload.Channels[2].Channel, 2)
	assert.Equal(t, payload.Payload.Channels[2].Type, "I")
	assert.EqualValues(t, payload.Payload.Channels[2].Value, 1)

	// int64
	payload.AddChannelByInt64(3, int64(1))

	assert.Len(t, payload.Payload.Channels, 4)
	assert.EqualValues(t, payload.Payload.Channels[3].Channel, 3)
	assert.Equal(t, payload.Payload.Channels[3].Type, "l")
	assert.EqualValues(t, payload.Payload.Channels[3].Value, 1)

	// uint64
	payload.AddChannelByUint64(4, uint64(1))

	assert.Len(t, payload.Payload.Channels, 5)
	assert.EqualValues(t, payload.Payload.Channels[4].Channel, 4)
	assert.Equal(t, payload.Payload.Channels[4].Type, "L")
	assert.EqualValues(t, payload.Payload.Channels[4].Value, 1)

	// float
	payload.AddChannelByFloat(5, float32(1))

	assert.Len(t, payload.Payload.Channels, 6)
	assert.EqualValues(t, payload.Payload.Channels[5].Channel, 5)
	assert.Equal(t, payload.Payload.Channels[5].Type, "f")
	assert.EqualValues(t, payload.Payload.Channels[5].Value, 1)

	// double
	payload.AddChannelByDouble(6, float64(1))

	assert.Len(t, payload.Payload.Channels, 7)
	assert.EqualValues(t, payload.Payload.Channels[6].Channel, 6)
	assert.Equal(t, payload.Payload.Channels[6].Type, "d")
	assert.EqualValues(t, payload.Payload.Channels[6].Value, 1)

	// clear channels
	payload.ClearChannels()
	assert.Len(t, payload.Payload.Channels, 0)
}
