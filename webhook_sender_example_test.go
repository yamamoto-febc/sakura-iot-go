package sakura_test

import (
	sakura "github.com/yamamoto-febc/sakura-iot-go"
)

func Example_send() {

	token := "[put your token]"
	secret := "[empty or put your secret]"
	module := "[put your module id]"

	// create sender
	sender := sakura.NewWebhookSender(token, secret)

	// create Payload
	p := sakura.NewPayload(module)

	p.AddValueByInt(0, int32(1))                 // ch:0 , set value(int32)
	p.AddValueByUint(1, uint32(1))               // ch:1 , set value(uint32)
	p.AddValueByInt64(2, int64(1))               // ch:2 , set value(int64)
	p.AddValueByUint64(3, uint64(1))             // ch:3 , set value(uint64)
	p.AddValueByFloat(4, float32(1))             // ch:4 , set value(float)
	p.AddValueByDouble(5, float64(1))            // ch:5 , set value(double)
	p.AddValueByHexString(6, "0f1e2d3c4b5c6b7a") // ch:6 , set value(HexString)

	err := sender.Send(p)
	if err != nil {
		panic(err)
	}

}
