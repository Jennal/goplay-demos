# How to run this Demo

## Init GoPlay Framework

```bash
go get github.com/jennal/goplay
```

## Server

```bash
go run chat/server/main.go
```

## Client

```bash
go run chat/client/main.go
```

Now every line you input to terminal will notify to server, and server will push back the same message with your name, and same message to every client which joined the same room.