package sakura_iot_go

import (
	"fmt"
	"time"
)

type Payload struct {
	Datetime time.Time    `json:"datetime"`
	Module   string       `json:"module"`
	Payload  InnerPayload `json:"payload"`
	Type     string       `json:"type"`
}

var (
	PayloadTypesKeepAlive = "keepalive"
	PayloadTypesChannels  = "channels"
)

func (p *Payload) IsKeepAlive() bool {
	return p.Type == PayloadTypesKeepAlive
}

func (p *Payload) IsChannelValue() bool {
	return p.Type == PayloadTypesChannels
}

type InnerPayload struct {
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Channel  int64       `json:"channel"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	Datetime time.Time   `json:"datetime"`
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

func (c *Channel) GetInt() (int32, error) {
	if c.Value == nil {
		return int32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return int32(v), nil
	}

	return int32(0), fmt.Errorf("Value is not a number")
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

func (c *Channel) GetInt64() (int64, error) {
	if c.Value == nil {
		return int64(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return int64(v), nil
	}

	return int64(0), fmt.Errorf("Value is not a number")
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

func (c *Channel) GetFloat() (float32, error) {
	if c.Value == nil {
		return float32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return float32(v), nil
	}

	return float32(0), fmt.Errorf("Value is not a number")
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
