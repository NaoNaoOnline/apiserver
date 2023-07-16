# apiserver

Golang based gRPC apiserver.


Running the apiserver locally.

```
./apiserver daemon
```

Calling the apiserver locally.

```
grpcurl -plaintext 127.0.0.1:7777 label.API/Search
```
