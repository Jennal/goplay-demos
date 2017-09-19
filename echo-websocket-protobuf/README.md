# How to run this Demo

## Init GoPlay Framework

```bash
go get -u github.com/jennal/goplay
```

## Server

```bash
go run echo-websocket-protobuf/server/main.go
```

## Go-Client

```bash
go run echo-websocket-protobuf/client/main.go
```

Now every line you input to terminal will notify & request to server, and server will push & response back the same line.

## Javascript-Client

See [goplay-client-javascript](https://github.com/Jennal/goplay-client-javascript/blob/master/demo/echo-websocket-protobuf/index.html)