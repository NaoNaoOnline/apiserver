# apiserver

Golang based [Twirp] apiserver. Requires at least `go1.21.1`.



Running redis stack locally.

```
docker run --rm --name redis-stack-apiserver -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```



Running the apiserver locally.

```
./apiserver daemon
```

```
{ "time":"2023-10-02 21:08:10", "leve":"info", "mess":"worker searching for tasks", "addr":"127.0.0.1:6379", "call":"/Users/xh3b4sd/project/NaoNaoOnline/apiserver/pkg/worker/worker.go:50" }
{ "time":"2023-10-02 21:08:10", "leve":"info", "mess":"server listening for calls", "addr":"127.0.0.1:7777", "call":"/Users/xh3b4sd/project/NaoNaoOnline/apiserver/pkg/server/server.go:69" }
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



Generating smart contract bindings using [abigen].

```
abigen --abi pkg/contract/policycontract/Policy.ABI.json --pkg policycontract --type Policy --out pkg/contract/policycontract/policy_contract.go
```



[abigen]: https://geth.ethereum.org/docs/tools/abigen
[Twirp]: https://github.com/twitchtv/twirp
