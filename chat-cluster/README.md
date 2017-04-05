# How to run this Demo

## Init GoPlay Framework

```bash
go get github.com/jennal/goplay
```

## Server

Before you run this server, you need to get [master server](https://github.com/Jennal/goplay-master) and [connector server](https://github.com/Jennal/goplay-connector) started.

### Start master server

```bash
go install github.com/jennal/goplay-master
$GOPATH/bin/goplay-master
```

### Start connector server

```bash
go install github.com/jennal/goplay-connector
$GOPATH/bin/goplay-connector
```

### Start chat-cluster server

```bash
go run chat-cluster/server/main.go
```

## Client

```bash
go run chat-cluster/client/main.go
```

Now every line you input to terminal will notify & request to server, and server will push & response back the same line.

You may notice that all the code in `chat/client/main.go` and `chat-cluster/client/main.go` are almost same. The only different is the port client connected. `chat/client/main.go` connect to chat server directly, and `chat-cluster/client/main.go` connect to `connector server` instead.