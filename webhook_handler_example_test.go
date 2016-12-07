package sakura_iot_go_test

import (
	"fmt"
	sakura "github.com/yamamoto-febc/sakura-iot-go"
	"net/http"
)

func Example_receive() {

	// Listen
	http.Handle("/", &sakura.WebhookHandler{
		Secret: "[put your secret]",
		HandleFunc: func(p sakura.Payload) {

			// [ここにWebhook 受信時の処理を書く]

			fmt.Printf("Module:%s\n", p.Module)
			fmt.Printf("Type  :%s\n", p.Type)
			fmt.Printf("Channels:%#v\n", p.Payload.Channels)

		},
	})
	http.ListenAndServe(":8080", nil)
}
