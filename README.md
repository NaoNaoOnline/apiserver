# apiserver

Golang based [Twirp] apiserver. Requires at least `go1.21.1`.



Running redis stack locally.

```
docker run --rm --name redis-stack-apiserver -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```



Filling the apiserver with test data.

```
./apiserver fakeit
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
curl -s --request "POST" --header "Content-Type: application/json" --data '{"object":[{"intern":{"user":"1699746343006008"}}]}' http://127.0.0.1:7777/event.API/Search | jq '.object[0]'
```

```
{
  "extern": [
    {
      "amnt": "43",
      "kind": "link",
      "user": false
    }
  ],
  "intern": {
    "crtd": "1699799790",
    "evnt": "1699799790111846",
    "user": "1699746343006008"
  },
  "public": {
    "cate": "1699746125560859,1699746343054161,1699746006668458,",
    "dura": "5400",
    "host": "1699746006088467,1699746346734161,1699746125979691",
    "link": "https://internationalinfomediaries.com.org",
    "time": "1699804800"
  }
}

```



Generating smart contract bindings using [abigen].

```
abigen --abi pkg/contract/policycontract/Policy.ABI.json --pkg policycontract --type Policy --out pkg/contract/policycontract/policy_contract.go
```



[abigen]: https://geth.ethereum.org/docs/tools/abigen
[Twirp]: https://github.com/twitchtv/twirp
