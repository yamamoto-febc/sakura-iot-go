# sakura-iot-go

[さくらのIoT Platform](https://iot.sakura.ad.jp)との連携サービス用ライブラリ

**This project is still developing.**

[![GoDoc](https://godoc.org/github.com/sacloud/libsacloud?status.svg)](https://godoc.org/github.com/yamamoto-febc/sakura-iot-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/sacloud/libsacloud)](https://goreportcard.com/report/github.com/yamamoto-febc/sakura-iot-go)


## 概要

`sakura-iot-go`は以下を提供しています。

- さくらのIoT Platformとの連携を行うためのライブラリ
  - HTTPハンドラ(net/http)
  - ペイロード用構造体の定義
  - Webhook送信(さくらのIoT Platform上の"Incoming Webhook"へのPOST)

- HTTPハンドラのサンプル実装としてエコーサーバー

HTTPハンドラはさくらのIoT PlatformからのOutgoing Webhookを受信し、

  - ペイロードの解析
  - HMAC-SHA1でのメッセージ署名検証
  
を行います。

## ライブラリとしての利用

#### `net/http`ライブラリでHTTPサーバーを起動する例

```golang

package main

import (
	"fmt"
	sakura "github.com/yamamoto-febc/sakura-iot-go"
	"net/http"
)

func main() {

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


```

#### さくらのIoT Platform上の"Incoming Webhook"へPOSTする例

```golang
package main

import (
	sakura "github.com/yamamoto-febc/sakura-iot-go"
)

func main() {

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

```

## サンプル実装(エコーサーバー) : Goビルド環境がある場合

```bash
# エコーサーバーの起動
$ go run cmd/echo_server.go

# 各種オプションの指定ありの場合
$ go run cmd/echo_server.go --port 8081 --path "/webhook" --secret "put your secret"  --debug

# ヘルプ:指定できるオプションの説明など
$ go run cmd/echo_server.go --help
```

## サンプル実装(エコーサーバー) : Goビルド環境がない場合

[リリースページ](https://github.com/yamamoto-febc/sakura-iot-go/releases/latest)にて実行ファイルを配布しています。

ダウンロードして展開、実行権を付与してください。
(以下の例ではカレントディレクトリに展開した場合のものです)

```bash
# エコーサーバーの起動
$ ./sakura-iot-echo-server

# 各種オプションの指定ありの場合
$ ./sakura-iot-echo-server --port 8081 --path "/webhook" --secret "put your secret" --debug

# ヘルプ:指定できるオプションの説明など
$ ./sakura-iot-echo-server --help
```

## License

 `sakura-iot-go` Copyright (C) 2016 Kazumichi Yamamoto.

  This project is published under [Apache 2.0 License](LICENSE.txt).
  
## Author

  * Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))

