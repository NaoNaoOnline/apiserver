# apiserver

Golang based [Twirp] apiserver.



Running the apiserver locally.

```
./apiserver daemon
```

```
{ "time":"2023-08-27 12:40:15", "leve":"info", "mess":"rpc server running at 127.0.0.1:7777", "call":"/Users/xh3b4sd/project/NaoNaoOnline/apiserver/pkg/server/server.go:66" }
```



Running redis stack locally.

```
docker run --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```



Calling the apiserver locally.

```
curl -s --request "POST" --header "Content-Type: application/json" --data '{}' http://127.0.0.1:7777/reaction.API/Search | jq '.object[0]'
```

```
{
  "intern": {
    "crtd": "1692392942",
    "rctn": "1692392942673667",
    "user": ""
  },
  "public": {
    "html": "️😍",
    "name": "smiling-face-with-heart-eyes"
  }
}
```



[Twirp]: https://github.com/twitchtv/twirp
