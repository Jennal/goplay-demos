# How to run this Demo

## Init GoPlay Framework

```bash
go get github.com/jennal/goplay
```

## Server

Before you run this server, you need to get [master server](https://github.com/Jennal/goplay-master).

### Start master server

```bash
go install github.com/jennal/goplay-master
$GOPATH/bin/goplay-master
```

### Start chat-cluster server

```bash
# Start 1st
go run chat-cluster/server/main.go
# Start 2nd
go run chat-cluster/server/main.go -p 2235
```

### Start connector server

If there are more than 1 Backend Server. [Connector server](https://github.com/Jennal/goplay-connector) should be started after backend.

```bash
go install github.com/jennal/goplay-connector

# Start 1st
$GOPATH/bin/goplay-connector
# Start 2nd
$GOPATH/bin/goplay-connector -p 9935
```

## Client

```bash
# Start 1st
go run chat-cluster/client/main.go
# Start 2nd
go run chat-cluster/client/main.go -p 9935
```

Now every line you input to terminal will notify to server, and server will push back the same line with name.

You may notice that all the code in `chat/client/main.go` and `chat-cluster/client/main.go` are almost same. The only different is the port client connected. `chat/client/main.go` connect to chat server directly, and `chat-cluster/client/main.go` connect to `connector server` instead.