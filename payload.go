package sakura

import (
	"fmt"
	"time"
)

var (
	// PayloadTypesKeepAlive WebSocket利用時のキープアライブを表すペイロードタイプ
	PayloadTypesKeepAlive = "keepalive"
	// PayloadTypesChannels データ送受信メッセージを表すペイロードタイプ
	PayloadTypesChannels = "channels"
	// PayloadTypesConnection モジュール接続時メッセージを表すペイロードタイプ
	PayloadTypesConnection = "connection"
)

// Payload Webhook/WebSocketでやりとりされるデータのペイロード
type Payload struct {
	Datetime *time.Time   `json:"datetime,omitempty"`
	Module   string       `json:"module"`
	Payload  InnerPayload `json:"payload"`
	Type     string       `json:"type"`
}

// NewPayload 新規ペイロード作成
func NewPayload(module string) Payload {
	return Payload{
		Module: module,
		Type:   PayloadTypesChannels,
		Payload: InnerPayload{
			Channels: []Channel{},
		},
	}
}

// IsKeepAlive ペイロードタイプがキープアライブであるか判定
func (p *Payload) IsKeepAlive() bool {
	return p.Type == PayloadTypesKeepAlive
}

// IsChannelValue ペイロードタイプがデータ送受信メッセージであるか判定
func (p *Payload) IsChannelValue() bool {
	return p.Type == PayloadTypesChannels
}

// IsConnection ペイロードタイプがモジュール接続時メッセージであるか確認
func (p *Payload) IsConnection() bool {
	return p.Type == PayloadTypesConnection
}

func (p *Payload) addChannel(c Channel) {
	p.Payload.Channels = append(p.Payload.Channels, c)
}

// AddValueByHexString 16進文字列(16文字で1セット)型の値を指定チャンネルに追加
func (p *Payload) AddValueByHexString(channel int64, value string) {
	c := newChannel(channel)
	c.SetHexString(value)
	p.addChannel(c)
}

// AddValueByInt int32型の値を指定チャンネルに追加
func (p *Payload) AddValueByInt(channel int64, value int32) {
	c := newChannel(channel)
	c.SetInt(value)
	p.addChannel(c)
}

// AddValueByUint uint32型の値を指定チャンネルに追加
func (p *Payload) AddValueByUint(channel int64, value uint32) {
	c := newChannel(channel)
	c.SetUint(value)
	p.addChannel(c)
}

// AddValueByInt64 int64型の値を指定チャンネルに追加
func (p *Payload) AddValueByInt64(channel int64, value int64) {
	c := newChannel(channel)
	c.SetInt64(value)
	p.addChannel(c)
}

// AddValueByUint64 uint64型の値を指定チャンネルに追加
func (p *Payload) AddValueByUint64(channel int64, value uint64) {
	c := newChannel(channel)
	c.SetUint64(value)
	p.addChannel(c)
}

// AddValueByFloat float(float32)型の値を指定チャンネルに追加
func (p *Payload) AddValueByFloat(channel int64, value float32) {
	c := newChannel(channel)
	c.SetFloat(value)
	p.addChannel(c)
}

// AddValueByDouble double(float64)型の値を指定チャンネルに追加
func (p *Payload) AddValueByDouble(channel int64, value float64) {
	c := newChannel(channel)
	c.SetDouble(value)
	p.addChannel(c)
}

// ClearValues ペイロードに含まれる全ての値をクリア
func (p *Payload) ClearValues() {
	p.Payload.Channels = []Channel{}
}

// ==============================================

// InnerPayload Payload内部の実データ格納用構造体リスト
type InnerPayload struct {
	Channels []Channel `json:"channels"`
}

// Channel Payload内部の実データ格納用の構造体
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

// GetHexString 16進文字列(16文字で1セット)を取得
func (c *Channel) GetHexString() (string, error) {
	if c.Value == nil {
		return "", fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(string); ok {
		return v, nil
	}

	return "", fmt.Errorf("Value is not HexString")
}

// SetHexString 16進文字列(16文字で1セット)を設定
func (c *Channel) SetHexString(v string) {
	c.Value = v
	c.Type = "b"
}

// GetInt int32型データを取得
func (c *Channel) GetInt() (int32, error) {
	if c.Value == nil {
		return int32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return int32(v), nil
	}

	return int32(0), fmt.Errorf("Value is not a number")
}

// SetInt int32型データを設定
func (c *Channel) SetInt(v int32) {
	c.Value = v
	c.Type = "i"
}

// GetUint uint32型データを取得
func (c *Channel) GetUint() (uint32, error) {
	if c.Value == nil {
		return uint32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return uint32(v), nil
	}

	return uint32(0), fmt.Errorf("Value is not a number")
}

// SetUint uint32型データを設定
func (c *Channel) SetUint(v uint32) {
	c.Value = v
	c.Type = "I"
}

// GetInt64 int64型データを取得
func (c *Channel) GetInt64() (int64, error) {
	if c.Value == nil {
		return int64(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return int64(v), nil
	}

	return int64(0), fmt.Errorf("Value is not a number")
}

// SetInt64 int64型データを設定
func (c *Channel) SetInt64(v int64) {
	c.Value = v
	c.Type = "l"
}

// GetUint64 uint64型データを取得
func (c *Channel) GetUint64() (uint64, error) {
	if c.Value == nil {
		return uint64(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return uint64(v), nil
	}

	return uint64(0), fmt.Errorf("Value is not a number")
}

// SetUint64 uint64型データを設定
func (c *Channel) SetUint64(v uint64) {
	c.Value = v
	c.Type = "L"
}

// GetFloat float(float32)型データを取得
func (c *Channel) GetFloat() (float32, error) {
	if c.Value == nil {
		return float32(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return float32(v), nil
	}

	return float32(0), fmt.Errorf("Value is not a number")
}

// SetFloat float(float32)型データを設定
func (c *Channel) SetFloat(v float32) {
	c.Value = v
	c.Type = "f"
}

// GetDouble double(float64)型データを取得
func (c *Channel) GetDouble() (float64, error) {
	if c.Value == nil {
		return float64(0), fmt.Errorf("Value is nil")
	}

	if v, ok := c.Value.(float64); ok {
		return float64(v), nil
	}

	return float64(0), fmt.Errorf("Value is not a number")
}

// SetDouble double(float64)型データを設定
func (c *Channel) SetDouble(v float64) {
	c.Value = v
	c.Type = "d"
}
