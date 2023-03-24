## Rate-Limite-Middleware
APIのレートリミットを制御するGoのecho用ミドルウェアです．
## Getting Started
1. `cd example`
2. `source .env` 環境変数の読み込み
3. `go run main.go`
4. `docker compose up -d --build` containerを立ち上げる
5. `docker logs curl --follow` curl側でリクエストが送信されていることをかくにんする
6. `docker logs nginx --follow` nginx側でのリクエスト処理を確認

## 動作の様子
![代替テキスト](img\play.gif)