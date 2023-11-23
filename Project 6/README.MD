## Client and server example with github.com/gorilla/websocket

This example shows a simple client and server.

The server echoes messages sent to it. The client sends a message every second and prints all messages received.

#### Setup:

```
go mod init websocket-example

go get github.com/gorilla/websocket
```

#### To run the example, start the server::

```
go run server.go
```

#### To run the example, start the server::

```
go run client.go
```

The server includes a simple web client. To use the client, open http://127.0.0.1:8080 in the browser and follow the instructions on the page.
