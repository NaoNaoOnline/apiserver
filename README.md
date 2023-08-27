# apiserver

Golang based [Twirp] apiserver.



Running the apiserver locally.

```
$ ./apiserver daemon
{ "tim":"2023-07-27 14:57:22", "lev":"inf", "mes":"rpc server running at 127.0.0.1:7777", "cal":"github.com/NaoNaoOnline/apiserver/cmd/daemon/run.go:99" }
```



Running redis stack locally.

```
docker run --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```



Calling the apiserver locally.

```
% curl -s --request "POST" --header "Content-Type: application/json" --data '{}' http://127.0.0.1:7777/reaction.API/Search | jq .
{
  "filter": null,
  "object": [
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
    },
    {
      "intern": {
        "crtd": "1692393021",
        "rctn": "1692393021407686",
        "user": ""
      },
      "public": {
        "html": "😂",
        "name": "face-with-tears-of-joy"
      }
    },
...
```



[Twirp]: https://github.com/twitchtv/twirp
