package sakura_iot_go

import (
	"fmt"
	"time"
)

var (
	PayloadTypesKeepAlive  = "keepalive"
	PayloadTypesChannels   = "channels"
	PayloadTypesConnection = "connection"
)

type Payload struct {
	Datetime *time.Time   `json:"datetime,omitempty"`
	Module   string       `json:"module"`
	Payload  InnerPayload `json:"payload"`
	Type     string       `json:"type"`
}

func NewPayload(module string) Payload {
	return Payload{
		Module: module,
		Type:   PayloadTypesChannels,
		Payload: InnerPayload{
			Channels: []Channel{},
		},
	}
}

func (p *Payload) IsKeepAlive() bool {
	return p.Type == PayloadTypesKeepAlive
}

func (p *Payload) IsChannelValue() bool {
	return p.Type == PayloadTypesChannels
}

func (p *Payload) IsConnection() bool {
	return p.Type == PayloadTypesConnection
}

func (p *Payload) addChannel(c Channel) {
	p.Payload.Channels = append(p.Payload.Channels, c)
}

func (p *Payload) AddChannelByHexString(channel int64, value string) {
	c := newChannel(channel)
	c.SetHexString(value)
	p.addChannel(c)
}

func (p *Payload) AddChannelByInt(channel int64, value int32) {
	c := newChannel(channel)
	c.SetInt(value)
	p.addChannel(c)
}

func (p *Payload) AddChannelByUint(channel int64, value uint32) {
	c := newChannel(channel)
	c.SetUint(value)
	p.addChannel(c)
}

func (p *Payload) AddChannelByInt64(channel int64, value int64) {
	c := newChannel(channel)
	c.SetInt64(value)
	p.addChannel(c)
}

func (p *Payload) AddChannelByUint64(channel int64, value uint64) {
	c := newChannel(channel)
	c.SetUint64(value)
	p.addChannel(c)
}

func (p *Payload) AddChannelByFloat(channel int64, value float32) {
	c := newChannel(channel)
	c.SetFloat(value)
	p.addChannel(c)
}

func (p *Payload) AddChannelByDouble(channel int64, value float64) {
	c := newChannel(channel)
	c.SetDouble(value)
	p.addChannel(c)
}

func (p *Payload) ClearChannels() {
	p.Payload.Channels = []Channel{}
}

// ==============================================

type InnerPayload struct {
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Channel  int64       `json:"channel"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	Datetime *time.Time  `json:"datetime,omitempty"`
}

func newChannel(channel int64) Channel {
	return Channel{
		Channel: channel,
	}
}

func (c *Channel) GetHexString() (string, error) {
	if c.Value == nil {
		return "", fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(string); ok {
		return v, nil
	}

	return "", fmt.Errorf("Value is not HexString")
}

func (c *Channel) SetHexString(v string) {
	c.Value = v
	c.Type = "b"
}

func (c *Channel) GetInt() (int32, error) {
	if c.Value == nil {
		return int32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return int32(v), nil
	}

	return int32(0), fmt.Errorf("Value is not a number")
}

func (c *Channel) SetInt(v int32) {
	c.Value = v
	c.Type = "i"
}

func (c *Channel) GetUint() (uint32, error) {
	if c.Value == nil {
		return uint32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return uint32(v), nil
	}

	return uint32(0), fmt.Errorf("Value is not a number")
}

func (c *Channel) SetUint(v uint32) {
	c.Value = v
	c.Type = "I"
}

func (c *Channel) GetInt64() (int64, error) {
	if c.Value == nil {
		return int64(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return int64(v), nil
	}

	return int64(0), fmt.Errorf("Value is not a number")
}

func (c *Channel) SetInt64(v int64) {
	c.Value = v
	c.Type = "l"
}

func (c *Channel) GetUint64() (uint64, error) {
	if c.Value == nil {
		return uint64(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return uint64(v), nil
	}

	return uint64(0), fmt.Errorf("Value is not a number")
}

func (c *Channel) SetUint64(v uint64) {
	c.Value = v
	c.Type = "L"
}

func (c *Channel) GetFloat() (float32, error) {
	if c.Value == nil {
		return float32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return float32(v), nil
	}

	return float32(0), fmt.Errorf("Value is not a number")
}

func (c *Channel) SetFloat(v float32) {
	c.Value = v
	c.Type = "f"
}

func (c *Channel) GetDouble() (float64, error) {
	if c.Value == nil {
		return float64(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return float64(v), nil
	}

	return float64(0), fmt.Errorf("Value is not a number")
}

func (c *Channel) SetDouble(v float64) {
	c.Value = v
	c.Type = "d"
}
