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

### Start echo-cluster server

```bash
go run echo-cluster/server/main.go
```

## Client

```bash
go run echo-cluster/client/main.go
```

Now every line you input to terminal will notify & request to server, and server will push & response back the same line.

You may notice that all the code in `echo/client/main.go` and `echo-cluster/client/main.go` are almost same. The only different is the port client connected. `echo/client/main.go` connect to echo server directly, and `echo-cluster/client/main.go` connect to `connector server` instead.